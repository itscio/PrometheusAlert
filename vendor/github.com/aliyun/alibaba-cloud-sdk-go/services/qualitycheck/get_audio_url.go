package qualitycheck

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetAudioUrl invokes the qualitycheck.GetAudioUrl API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/getaudiourl.html
func (client *Client) GetAudioUrl(request *GetAudioUrlRequest) (response *GetAudioUrlResponse, err error) {
	response = CreateGetAudioUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetAudioUrlWithChan invokes the qualitycheck.GetAudioUrl API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getaudiourl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAudioUrlWithChan(request *GetAudioUrlRequest) (<-chan *GetAudioUrlResponse, <-chan error) {
	responseChan := make(chan *GetAudioUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAudioUrl(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetAudioUrlWithCallback invokes the qualitycheck.GetAudioUrl API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getaudiourl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAudioUrlWithCallback(request *GetAudioUrlRequest, callback func(response *GetAudioUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAudioUrlResponse
		var err error
		defer close(result)
		response, err = client.GetAudioUrl(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetAudioUrlRequest is the request struct for api GetAudioUrl
type GetAudioUrlRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetAudioUrlResponse is the response struct for api GetAudioUrl
type GetAudioUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAudioUrlRequest creates a request to invoke GetAudioUrl API
func CreateGetAudioUrlRequest() (request *GetAudioUrlRequest) {
	request = &GetAudioUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetAudioUrl", "", "")
	return
}

// CreateGetAudioUrlResponse creates a response to parse from GetAudioUrl response
func CreateGetAudioUrlResponse() (response *GetAudioUrlResponse) {
	response = &GetAudioUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}