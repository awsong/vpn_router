// Nodejs encryption with GCM
// Does not work with nodejs v0.10.31
// Part of https://github.com/chris-rock/node-crypto-examples

var crypto = require('crypto'),
//  algorithm = 'aes-256-gcm',
  algorithm = 'aes-256-gcm',
  password = '3zTvzr3p67VC61jmV54rIYu1545x4TlY',
  // do not use a global iv for production, 
  // generate a new one for each encryption
  iv = crypto.randomBytes(12)

function encrypt(text) {
  var cipher = crypto.createCipheriv(algorithm, password, iv)
  var encrypted = cipher.update(text, 'utf8', 'hex')
  encrypted += cipher.final('hex');
  var tag = cipher.getAuthTag();
  return {
    content: encrypted,
    tag: tag
  };
}

var hw = encrypt("?a=100.22&p=ghost002&s=202.5.16.100")
var text = hw.content+hw.tag.toString('hex');
console.log(text);
var ivt =iv.toString('hex');
console.log(ivt);
  // outputs hello world
console.log(decrypt(hw));
urlString="curl \"http://localhost:29080/search?q=" + text + "&n=" + ivt + "\"";
console.log(urlString)
