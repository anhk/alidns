package main

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"net"
)

var (
	domainName = "anhk.cc"
	keyWord    = "files"
	infName    = "eno1"
)

func getIp(client *alidns.Client) (string, string, error) {
	req := alidns.CreateDescribeDomainRecordsRequest()
	req.DomainName = domainName
	req.KeyWord = keyWord

	res, err := client.DescribeDomainRecords(req)
	if err != nil {
		return "", "", err
	}

	for _, v := range res.DomainRecords.Record {
		return v.RecordId, v.Value, nil
	}
	return "", "", errors.New("no record.")
}

func setIp(client *alidns.Client, id, ip string) error {
	req := alidns.CreateUpdateDomainRecordRequest()
	req.RecordId = id
	req.Value = ip
	req.RR = keyWord
	req.Type = "A"

	_, err := client.UpdateDomainRecord(req)
	if err != nil {
		return err
	}
	return nil
}

func getSysIp() (string, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, v := range ifs {
		if v.Name != infName {
			continue
		}
		addrs, err := v.Addrs()
		if err != nil {
			return "", err
		}

		for _, vv := range addrs {
			ip, _, err := net.ParseCIDR(vv.String())
			if err != nil {
				return "", err
			}
			return ip.String(), nil
		}

	}
	return "", errors.New("no record.")
}

func main() {
	client, err := alidns.NewClientWithAccessKey("default", "<ak>", "<sk>")
	if err != nil {
		panic(err)
	}

	sysIp, err := getSysIp()
	if err != nil {
		panic(err)
	}

	id, ip, err := getIp(client)
	if err != nil {
		panic(err)
	}

	if sysIp == ip {
		return
	}

	if err := setIp(client, id, sysIp); err != nil {
		panic(err)
	}
}
