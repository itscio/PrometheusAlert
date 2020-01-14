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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// StartSynchronizationJob invokes the dts.StartSynchronizationJob API synchronously
// api document: https://help.aliyun.com/api/dts/startsynchronizationjob.html
func (client *Client) StartSynchronizationJob(request *StartSynchronizationJobRequest) (response *StartSynchronizationJobResponse, err error) {
	response = CreateStartSynchronizationJobResponse()
	err = client.DoAction(request, response)
	return
}

// StartSynchronizationJobWithChan invokes the dts.StartSynchronizationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/startsynchronizationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartSynchronizationJobWithChan(request *StartSynchronizationJobRequest) (<-chan *StartSynchronizationJobResponse, <-chan error) {
	responseChan := make(chan *StartSynchronizationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartSynchronizationJob(request)
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

// StartSynchronizationJobWithCallback invokes the dts.StartSynchronizationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/startsynchronizationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartSynchronizationJobWithCallback(request *StartSynchronizationJobRequest, callback func(response *StartSynchronizationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartSynchronizationJobResponse
		var err error
		defer close(result)
		response, err = client.StartSynchronizationJob(request)
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

// StartSynchronizationJobRequest is the request struct for api StartSynchronizationJob
type StartSynchronizationJobRequest struct {
	*requests.RpcRequest
	SynchronizationJobId     string `position:"Query" name:"SynchronizationJobId"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
}

// StartSynchronizationJobResponse is the response struct for api StartSynchronizationJob
type StartSynchronizationJobResponse struct {
	*responses.BaseResponse
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateStartSynchronizationJobRequest creates a request to invoke StartSynchronizationJob API
func CreateStartSynchronizationJobRequest() (request *StartSynchronizationJobRequest) {
	request = &StartSynchronizationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "StartSynchronizationJob", "dts", "openAPI")
	return
}

// CreateStartSynchronizationJobResponse creates a response to parse from StartSynchronizationJob response
func CreateStartSynchronizationJobResponse() (response *StartSynchronizationJobResponse) {
	response = &StartSynchronizationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}