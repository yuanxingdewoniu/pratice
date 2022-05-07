#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import jinja2
import tempfile
import logging
from pdfkit import from_string
from sys import argv, exit, stdin
from subprocess import check_output, run , Popen , PIPE


class Arg:
    def __init__(self, security, usename, filename):
        self.username = username
        self.filename = filename
        self.security = security


def getSecurity(filename):

    logging.debug("3")
#    try:
#        output = check_output(['touch', filename]).decode()
#    except subprocess.CalledProcessError as e:
#        output = e.output
#        code = e.returncode
#        logging.debug("error:{}".format(output))

#    logging.debug("4")
#    security = output.split(' ')[0].split(':')[-1]
    p = Popen(['ls','-Z'],stdout = PIPE, stdin = PIPE)
    stdout,stderr = p.communicate(b'''filename''')
    output = stdout.decode()
    security = output.split(' ')[0]
#    security = 's0'
    logging.debug("security:{}".format(security))

    return (security)


def generateps_in(data):
    logging.debug("1")
    with open('ps.in', 'w') as f:
        f.write(data.decode())
    logging.debug("2")


if __name__ == '__main__':
    logging.basicConfig(filename='/tmp/watermarkpdf.log', level=logging.DEBUG)

    logging.debug("argv: {}".format(argv))

    template_search_path = '/usr/share/deepin-security/'
    css_file = template_search_path + '/style.css'

    tempdir = tempfile.mkdtemp()
    logging.debug("tempdir: {}".format(tempdir))

    os.chdir(tempdir)

    username = argv[2]
    logging.debug("username: {}".format(username))
#    filename = argv[5].split()[-1].split('=')[-1]
    filename = argv[3]
    logging.debug("filename: {}".format(filename))

    try:
        if len(argv) >= 7:
            generateps_in(argv[6])
        else:
            generateps_in(stdin.buffer.read())
    except Exception as e:
        logging.debug("error: {}".format(e))
        exit(1)

    security = getSecurity(filename)
    logging.debug("security: {}".format(security))
    args = Arg(security, username, filename)
    logging.debug("args: {}".format(args))

    loader = jinja2.FileSystemLoader(searchpath=template_search_path)
    environment = jinja2.Environment(loader=loader)
    template = environment.get_template('template.html')
    outText = template.render(args=args)
    logging.debug("outText: {}".format(outText))

    logging.debug("cwd: {}".format(os.getcwd()))
    try:
        from_string(outText, 'out.pdf', options={'quiet': ''}, css=css_file)
    except Exception as e:
        logging.debug("error: {}".format(e))
    logging.debug("produce out.pdf successfully")

    run(['/usr/bin/ps2pdf14', 'ps.in', 'pdf.in'])
    logging.debug("ps2pdf14")
    run(['/usr/bin/pdftk', 'pdf.in', 'background', 'out.pdf',
         'output', 'pdf.out'])
    logging.debug("pdftk")
    run(['/usr/bin/pdftops', 'pdf.out', 'ps.out'])
    logging.debug("pdftops")

    with open('ps.out', 'r') as f:
        print(f.read())
