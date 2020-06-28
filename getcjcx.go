/*
 * Copyright 2020 CDFMLR
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package qzgo

import "log"

/*
 GetCjcx 获取成绩信息

 GET http://jwxt.xxxx.edu.cn/app.do?method=getCjcx&xh={$学号}&xnxqid={$学年学期ID}

 Req:
	request.header{token:'运行身份验证authUser时获取到的token，有过期机制'},
	request.data{
		'method':'getCjcx',  //必填
		'xh':'2017168xxxxx',  //必填，可以添加非本token学号查询他人成绩
		'xqxnid':'2017-2018-2' //非必填，不填输出全部成绩
	}

 Resp:
	{
		"success": ture,
		"result": 	[
			{
				"bz": null,  //未知
				"cjbsmc": null,  //特殊情况通报，例如“作弊”“缺考”
				"kclbmc": "必修", //课程类别名称
				"zcj": "88",  //总成绩
				"xm": "某某某",  //学生姓名
				"xqmc": "2017-2018-2",  //学期名称
				"kcxzmc": "公共基础课",  //课程性质名称，根据此项不同值可判断该科成绩是否计入GPA
				"kcywmc": "College students career development and guidance",  //课程英文名称
				"ksxzmc": "正常考试",  //考试性质名称,目前遇见的情况有正常考试，补考（x），重修（x），分别意为补考第x次和重修第x次，若补考未通过，正常考试条目和补考条目将同时存在，若补考通过，则只存在补考条目
				"kcmc": "大学生职业发展与就业指导",  //课程名称
				"xf": 1  //学分
			}
		]
	}

*/
func (c *Client) GetCjcx(xh, xnxqid string) (GetCjcxRespBody, error) {
	var resp GetCjcxRespBody
	q := map[string]string{
		"method": "getCjcx",
		"xh":     xh,
		"xnxqid": xnxqid,
	}
	err := qzApiGet(c.School, c.token, &resp, q)
	if err != nil {
		log.Println(err)
		return GetCjcxRespBody{}, err
	}
	return resp, nil
}

type GetCjcxRespBody struct {
	Success bool                  `json:"success"`
	Result  []GetCjcxRespBodyItem `json:"result"`
}

type GetCjcxRespBodyItem struct {
	Cjbsmc string  `json:"cjbsmc"`
	Kclbmc string  `json:"kclbmc"`
	Zcj    string  `json:"zcj"`
	Xm     string  `json:"xm"`
	Xqmc   string  `json:"xqmc"`
	Kcxzmc string  `json:"kcxzmc"`
	Kcywmc string  `json:"kcywmc"`
	Ksxzmc string  `json:"ksxzmc"`
	Kcmc   string  `json:"kcmc"`
	Xf     float32 `json:"xf"`
}
