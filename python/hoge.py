from __future__ import division, print_function, absolute_import, unicode_literals
#import sys, subprocess, os.path, mysql.connector
#import glob


class runtime:
    @staticmethod
    def v3():
        return sys.version_info >= (3,)

    @staticmethod
    def v2():
        return sys.version_info < (3,)


class colors:
    bold = '\033[1m'
    underlined = '\033[4m'

    black = '\033[30m'
    red = '\033[31m'
    green = '\033[32m'
    yellow = '\033[33m'
    blue = '\033[34m'
    magenta = '\033[35m'
    cyan = '\033[36m'
    lightgray = '\033[37m'
    darkgray = '\033[90m'
    lightred = '\033[91m'
    lightgreen = '\033[92m'
    lightyellow = '\033[93m'
    lightblue = '\033[94m'
    lightmagenta = '\033[95m'
    lightcyan = '\033[96m'
    
    background_black = '\033[40m'
    background_red = '\033[41m'
    background_green = '\033[42m'
    background_yellow = '\033[43m'
    background_blue = '\033[44m'
    background_magenta = '\033[45m'
    background_cyan = '\033[46m'

    reset = '\033[0m'


# class colors
for i in range(1, 10):
   print(colors.red + str(i) + colors.reset)


# execute shell command
#out = subprocess.check_output("ls -la",  shell=True).strip()
#if runtime.v3(): out = out.decode('utf-8')
#print(out)


# file list
#print(glob.glob("/tmp/*"))
#path = '/tmp'
#for root, dirs, files in os.walk(path):
#    for file in files:
#        print(os.path.join(root, file))


# execute mysql
#if __name__ == '__main__':
#
#    connect = mysql.connector.connect(user='root', password='', host='localhost', database='sample', charset='utf8')
#    cursor = connect.cursor()
#
##    name = 'yamada taro'
##    sex = 'man'
##
##    # insert
##    cursor.execute('insert into student_table (name, sex) values (%s, %s)', (name, sex))
#
#    # select
#    cursor.execute('select id from student_table')
##	row = cursor.fetchone()
#    row = cursor.fetchall()
#
#    # print
#    for i in row:
#        print(i[0])
#
#    # db connection close
#	cursor.close()
#	connect.close()
