import os
import yaml
import string
import codecs

def main():

    # template を開く
    tmpf = open("template.html", "r")
    template = string.Template(tmpf.read())

    # yaml から値を取得する
    f = open("config.yml", "r")
    data = yaml.load(f)

    # yaml の設定分ループ
    for i in data:
        # yaml の値で、template を書き換える
        conf = data[i]
        value = template.safe_substitute(conf)

        # ./html/ に書き換えた template を保存
        name = "./html/" + i + ".html"
        scriptf = codecs.open(name, "w", 'utf-8')
        scriptf.write(value)
        scriptf.close()

    tmpf.close()

if __name__ == '__main__':
    main()
