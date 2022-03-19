//加载nodejs的模块，模块名叫ethereumjs-wallet
var Wallet = require('ethereumjs-wallet');
//填入自己的密钥
var key = Buffer.from('cbce2b021617132efec8eb565d12ff7451d49c00ede666e8ed3f099cad6810c2', 'hex');
var wallet = Wallet.fromPrivateKey(key);
//填入自己设置的密码
var keystore = wallet.toV3String('20201004');

console.log(keystore);
