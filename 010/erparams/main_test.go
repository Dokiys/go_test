package erparams

import (
	"testing"
)

func TestAll(t *testing.T) {
	str := "[{\"key\":\"Row0\",\"api\":\"http://api.easyreport.fancydsp.com/api/projects/4/queries/1349/results\",\"param\":\"{\\\"start_date\\\":\\\"$today\\\",\\\"hour\\\":\\\"$thishour\\\",\\\"final_id\\\":\\\"\\\",\\\"date_form\\\":\\\"%Y/%m/%d\\\",\\\"date_before\\\":\\\"\\\",\\\"date_after\\\":\\\"\\\",\\\"hour_form\\\":\\\":00\\\",\\\"dim\\\":\\\"cost,y_all_cost\\\",\\\"map\\\":\\\"[{\\\\\\\"mapping_field\\\\\\\":\\\\\\\"vendor_advertiser_id\\\\\\\",\\\\\\\"mapping_name\\\\\\\":\\\\\\\"sgxwcvfq\\\\\\\",\\\\\\\"default_value\\\\\\\":\\\\\\\"-\\\\\\\",\\\\\\\"mapping\\\\\\\":{\\\\\\\"10474236\\\\\\\":\\\\\\\"User_1582630802878\\\\\\\",\\\\\\\"10474238\\\\\\\":\\\\\\\"鲜衣怒马-6\\\\\\\",\\\\\\\"10474240\\\\\\\":\\\\\\\"测试账户1\\\\\\\"}}]\\\",\\\"const\\\":\\\"[{\\\\\\\"custom_field\\\\\\\":\\\\\\\"cqnpqolr\\\\\\\",\\\\\\\"expression\\\\\\\":{\\\\\\\"op\\\\\\\":\\\\\\\"cst\\\\\\\",\\\\\\\"params\\\\\\\":\\\\\\\"'-'\\\\\\\"}},{\\\\\\\"custom_field\\\\\\\":\\\\\\\"wxouuwnf\\\\\\\",\\\\\\\"expression\\\\\\\":{\\\\\\\"op\\\\\\\":\\\\\\\"cst\\\\\\\",\\\\\\\"params\\\\\\\":\\\\\\\"'-'\\\\\\\"}},{\\\\\\\"custom_field\\\\\\\":\\\\\\\"bezamkfp\\\\\\\",\\\\\\\"expression\\\\\\\":{\\\\\\\"op\\\\\\\":\\\\\\\"cst\\\\\\\",\\\\\\\"params\\\\\\\":\\\\\\\"'-'\\\\\\\"}}]\\\",\\\"cal\\\":\\\"[]\\\",\\\"group\\\":\\\"vendor_advertiser_id order by toInt64(vendor_advertiser_id)\\\",\\\"_aggregate\\\":\\\"[{\\\\\\\"field_name\\\\\\\":\\\\\\\"cqnpqolr\\\\\\\",\\\\\\\"expression\\\\\\\":\\\\\\\"sum(cost)\\\\\\\",\\\\\\\"method\\\\\\\":\\\\\\\"exp\\\\\\\"},{\\\\\\\"field_name\\\\\\\":\\\\\\\"wxouuwnf\\\\\\\",\\\\\\\"expression\\\\\\\":\\\\\\\"sum(y_all_cost)\\\\\\\",\\\\\\\"method\\\\\\\":\\\\\\\"exp\\\\\\\"},{\\\\\\\"field_name\\\\\\\":\\\\\\\"bezamkfp\\\\\\\",\\\\\\\"expression\\\\\\\":\\\\\\\"sum(cost)/sum(y_all_cost)\\\\\\\",\\\\\\\"method\\\\\\\":\\\\\\\"exp\\\\\\\"}]\\\"}\"}]"
	sendConfigs,err := parseSendConfig(str)
	if err != nil {
		panic(err)
	}
	err = printAll(sendConfigs, "")
	if err != nil {
		panic(err)
	}
}