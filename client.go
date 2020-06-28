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

package qzapi

type Client struct {
	School string // 学校
	Xh     string // 学号
	Pwd    string // 密码

	token string
}

func NewClientLogin(school string, xh string, pwd string) (*Client, error) {
	client := &Client{School: school, Xh: xh, Pwd: pwd}
	_, err := client.AuthUser()
	return client, err
}

