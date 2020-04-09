/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type DpBlockBackupJobsApiService service


/* DpBlockBackupJobsApiService 
 List dp block backup jobs
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "limit" (int64) paging param
     @param "offset" (int64) paging param
     @param "q" (string) query param of search
     @param "sort" (string) sort param of search
     @param "blockVolumeId" (int64) block volume
     @param "blockSnapshotId" (int64) object storage bucket
     @param "dpBlockBackupPolicyId" (int64) dp block backup policy
 @return DpBlockBackupJobsResp*/
func (a *DpBlockBackupJobsApiService) ListDpBlockBackupJobs(ctx context.Context, localVarOptionals map[string]interface{}) (DpBlockBackupJobsResp,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DpBlockBackupJobsResp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/dp-block-backup-jobs/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["limit"], "int64", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["offset"], "int64", "offset"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["q"], "string", "q"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sort"], "string", "sort"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["blockVolumeId"], "int64", "blockVolumeId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["blockSnapshotId"], "int64", "blockSnapshotId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["dpBlockBackupPolicyId"], "int64", "dpBlockBackupPolicyId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int64); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["offset"].(int64); localVarOk {
		localVarQueryParams.Add("offset", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["q"].(string); localVarOk {
		localVarQueryParams.Add("q", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sort"].(string); localVarOk {
		localVarQueryParams.Add("sort", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["blockVolumeId"].(int64); localVarOk {
		localVarQueryParams.Add("block_volume_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["blockSnapshotId"].(int64); localVarOk {
		localVarQueryParams.Add("block_snapshot_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["dpBlockBackupPolicyId"].(int64); localVarOk {
		localVarQueryParams.Add("dp_block_backup_policy_id", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Xms-Auth-Token"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("token", key)
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

