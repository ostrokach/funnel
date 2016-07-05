#!/usr/bin/env python

import unittest
import uuid
import time
import requests

from common_test_util import ServerTest, get_abspath

class TestFileOP(ServerTest):


    def test_file_mount(self):
        
        loc = self.copy_to_storage( get_abspath("test_data.1") )

        task = {
            "name" : "TestMD5",
            "projectId" : "MyProject",
            "description" : "My Desc",
            "inputs" : [
                {
                    "name" : "infile",
                    "description" : "File to be MD5ed",
                    "location" : loc,
                    "path" : "/tmp/test_file"
                }
            ],
            "outputs" : [
                {
                    "location" : "test_data.out",
                    "path" : "/tmp/test_out"
                }
            ],
            "resources" : {
                "volumes" : [{
                    "name" : "test_disk",
                    "sizeGb" : 5,
                    "mountPoint" : "/tmp"
                }]
            },
            "docker" : [
                {
                    "imageName" : "ubuntu",
                    "cmd" : ["md5sum", "/tmp/test_file"],
                    "stdout" : "/tmp/test_out"
                }
            ]
        }

        r = requests.post("http://localhost:8000/v1/jobs", json=task)
        data = r.json()
        print data
        job_id = data['value']

        for i in range(10):
            r = requests.get("http://localhost:8000/v1/jobs/%s" % (job_id))
            data = r.json()
            if data["state"] not in ['Queued', "Running"]:
                break
            time.sleep(1)
        print data

        #assert 'logs' in data
        #assert data['logs'][0]['stdout'] == "hello world\n"


