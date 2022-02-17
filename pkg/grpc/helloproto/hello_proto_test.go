package helloproto

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"io/ioutil"
	"log"
	"testing"
)

func TestHelloProto(t *testing.T) {
	// The whole purpose of using protocol buffers is to serialize your data
	bookOut := &AddressBook{
		People: []*Person{{
			Id:    1234,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*Person_PhoneNumber{
				{Number: "555-4321", Type: Person_HOME},
			},
		}},
	}
	data, err := proto.Marshal(bookOut)
	assert.NoError(t, err)

	bookIn := &AddressBook{}
	assert.NoError(t, proto.Unmarshal(data, bookIn))
}

func TestHelloProtoReflect(t *testing.T) {
	book := &AddressBook{
		People: []*Person{{
			Id:    1234,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*Person_PhoneNumber{
				{Number: "555-4321", Type: Person_HOME},
			},
		}},
	}
	data, err := proto.Marshal(book)
	assert.NoError(t, err)

	// Get message type by full name
	msgType, err := protoregistry.GlobalTypes.FindMessageByName("helloproto.AddressBook")
	assert.NoError(t, err)

	// Deserialize into helloproto.AddressBook message
	msg := msgType.New().Interface()
	err = proto.Unmarshal(data, msg)
	assert.NoError(t, err)

	t.Log(msg)
}

func TestHelloProtoReflect2(t *testing.T) {
	person := &Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*Person_PhoneNumber{
			{Number: "555-4321", Type: Person_HOME},
		},
	}
	msg := person.ProtoReflect()
	msg.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Name() == "name" {
			msg.Set(fd, protoreflect.ValueOfString("zhangsan"))
		}
		return true
	})

	t.Log(msg.Interface())
}


func TestHelloDescProto(t *testing.T) {
	fds := &descriptorpb.FileDescriptorSet{}

	{
		descriptorFile := "./helloproto/descriptor.desc"
		// Open the descriptor
		b, err := ioutil.ReadFile(descriptorFile)
		assert.NoError(t, err)

		// Unmarshal descriptor file into the FileDescriptorSet
		err = proto.Unmarshal(b, fds)
		assert.NoError(t, err)
	}

	// Create a Files registry
	files, err := protodesc.NewFiles(fds)
	assert.NoError(t, err)

	// Query (Descriptors) by name
	d, err := files.FindDescriptorByName("helloproto.Person")
	assert.NoError(t, err)
	m, ok := d.(protoreflect.MessageDescriptor)
	if !ok {
		log.Fatal("Unable to assert into MessageDescriptor")
	}

	// Stores a value in a field
	person := dynamicpb.NewMessage(m)
	n, err := files.FindDescriptorByName("helloproto.Person.name")
	assert.NoError(t, err)
	f, ok := n.(protoreflect.FieldDescriptor)
	if !ok {
		log.Fatal("Unable to assert into FieldDescriptor")
	}
	person.Set(f, protoreflect.ValueOfString("zhangsan"))

	// Serialize
	b, err := proto.Marshal(person)
	assert.NoError(t, err)

	// Deserialize
	p := &Person{}
	err = proto.Unmarshal(b, p)
	assert.NoError(t, err)

	t.Log(p)
}
