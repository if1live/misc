#!/usr/bin/env python
#-*- coding: utf-8 -*-

import flask as fl
import subprocess

app = fl.Flask(__name__)

quiz_code = '''
static_assert(factorial<1>::value == 1, "");
static_assert(factorial<2>::value == 2, "");
static_assert(factorial<3>::value == 6, "");
static_assert(factorial<4>::value == 24, "");
'''

@app.route('/')
def index():
    return fl.render_template('index.jinja2', quiz_code=quiz_code)

@app.route('/submit', methods=['POST'])
def submit_code():
    code = fl.request.form['code']
    fullcode = '\n'.join([code, quiz_code])

    tmp_cpp_src = '/tmp/madoka-sample.cpp'
    f = open(tmp_cpp_src, 'wb')
    f.write(fullcode)
    f.close()

    s = subprocess.Popen(['clang', '-c', tmp_cpp_src, '-std=c++11'],
                         stdout=subprocess.PIPE,
                         stderr=subprocess.PIPE)

    stderr_list = []
    while True:
        line = s.stderr.readline()
        if not line:
            break
        stderr_list.append(line)

    success = (len(stderr_list) == 0)
    return fl.render_template('result.jinja2',
                              fullcode=fullcode,
                              stderr_list=stderr_list,
                              success=success)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
