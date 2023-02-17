import time

a = {'TaskId': '7c7b52f07b194f70b3c0404ba6fcc0f7', 'RequestId': 'EF0B3D65-0A39-5594-B80C-93C5FD8E45FD',
     'StatusText': 'SUCCESS', 'BizDuration': 3101, 'SolveTime': 1676555222573, 'RequestTime': 1676555218263,
     'StatusCode': 21050000, 'Result': {'Sentences': [
        {'EndTime': 2510, 'SilenceDuration': 0, 'BeginTime': 880, 'Text': '北京的天气。', 'ChannelISpeechRate': 184,
         'EmotionValue': 6.7}]}}
print(a)
print(a['TaskId'])
print(a['RequestId'])
print(a['StatusText'])
print(a['BizDuration'])
print(a['SolveTime'])
timeSolve = time.localtime(a['SolveTime'] / 1000)
SolveTime = time.strftime("%Y-%m-%d %H:%M:%S", timeSolve)
print(SolveTime)
print(a['RequestTime'])
RequestArray = time.localtime(a['RequestTime'] / 1000)
RequestTime = time.strftime("%Y-%m-%d %H:%M:%S", RequestArray)
print(RequestTime)
print(a['StatusCode'])
print(a['Result'])
print(a['Result']['Sentences'][0]['EndTime'])
print(a['Result']['Sentences'][0]['SilenceDuration'])
print(a['Result']['Sentences'][0]['BeginTime'])
print(a['Result']['Sentences'][0]['Text'])
print(a['Result']['Sentences'][0]['ChannelISpeechRate'])
print(a['Result']['Sentences'][0]['EmotionValue'])

writeHtml = f'''<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>VTOT</title>
</head>
<body>
<h1>VTOT</h1>
<pre>{a['Result']['Sentences'][0]['Text']}</pre>
<br>
<br>
<h6>开始时间：{RequestTime}</h6>
<h6>结束时间：{SolveTime}</h6>
</body>
</html>'''

print(writeHtml)
