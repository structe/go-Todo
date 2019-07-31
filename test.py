from selenium import webdriver
import requests

ip = "http://192.168.1.150:8000/meet/api"


'''
postdata={"email":"123@qq.com","password":"123"}
r=requests.post(ip+"/signup",json=postdata)
print(r.headers)
print(r.text)

postdata={"email":"123@qq.com","password":"123"}
r=requests.post(ip+"/signin",json=postdata)
print(r.headers)
print(r.text)

postdata={"userid":1,"timestamp":"10:23","content":"good",
                "status":0,"date":"2019-7-29"}
r=requests.post(ip+"/newtodo",json=postdata,cookies=cookies)
print(r.headers)
print(r.text)
'''

postdata={"email":"123@qq.com","password":"123"}
r=requests.post(ip+"/signin",json=postdata)
print(r.headers)
print(r.text)
cookies=dict(token=r.cookies['token'])

postdata={"userid":1,"timestamp":"12:23","content":"good",
                "status":0,"date":"2019-7-28"}
r=requests.post(ip+"/newtodo",json=postdata,cookies=cookies)
print(r.headers)
print(r.text)

postdata={"userid":1,"date":"2019-7-28"}
r=requests.post(ip+"/listtodo",json=postdata,cookies=cookies)
#r=requests.post(ip+"/listtodo",json=postdata)
print(r.headers)
print(r.text)




#driver =  webdriver.Firefox()
#response = driver.get("https://aliyun.com")
#print(response)  