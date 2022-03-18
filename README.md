# actions-test

GitHub Actionsで試してみたいことを実験する

## 知見

- envで先に定義した変数を，envで後に定義する変数に渡すことはできない(?)
- envで定義した変数 `XXX` をsteps内のenvで参照するには， `${{ env.XXX }}` とすれば良い
