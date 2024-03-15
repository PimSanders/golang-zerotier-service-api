package golangzerotierserviceapi

import (
	"bytes"
	"encoding/json"
)

func (c *Client) JoinNetwork(networkId string) (Network, error) {
	var respMap Network

	joinData := Join{
		AllowDNS:     true,
		AllowDefault: true,
		AllowManaged: true,
		AllowGlobal:  true,
	}

	dataJson, err := json.Marshal(joinData)
	if err != nil {
		return respMap, err
	}
	data := bytes.NewReader(dataJson)

	respBody, err := c.MakePostRequest("/network/"+networkId, data)

	if err != nil {
		return respMap, err
	}

	// Unmarshal the JSON response into the respMap struct
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		return respMap, err
	}

	return respMap, err

}

func (c *Client) LeaveNetwork(networkId string) (LeaveNetwork, error) {
	var respMap LeaveNetwork

	respBody, err := c.MakeDeleteRequest("/network/" + networkId)
	if err != nil {

		return respMap, err
	}

	// Unmarshal the JSON response into the respMap struct
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		return respMap, err
	}

	return respMap, err
}

func (c *Client) GetNetworkStatus() (Status, error) {
	var respMap Status
	respBody, err := c.MakeGetRequest("/status")

	if err != nil {

		return respMap, err
	}

	// Unmarshal the JSON response into the respMap struct
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		return respMap, err
	}

	return respMap, err
}
