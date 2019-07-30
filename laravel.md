## デプロイ時の最適化
```
composer install

// config/*.php の設定をキャッシュ
php artisan config:cache

// ルーティング設定をキャッシュ
php artisan route:cache
```


## composer install
composer.lock を元にパッケージをインストールし、
下記のコマンドが実行される

```
composer dump-autoload
php artisan clear-compiled
php artisan optimize
```


## キャッシュクリア
```
php artisan cache:clear
php artisan config:clear
php artisan route:clear
php artisan view:clear
```


## db 再作成 & seed 実行
```
php artisan migrate:refresh --seed
```


## seed 実行
```
php artisan db:seed
php artisan db:seed --class=SampleSeeder    // class 指定
```
