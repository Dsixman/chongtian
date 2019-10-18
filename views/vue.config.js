const webpack = require('webpack')
module.exports={
	devServer:{
		proxy:{
		'/api':{
			target:'http://192.168.0.107:9090/',
			ws:true,
			changeOrigin:true,
			pathRewrite:{
				'^/api':''
			}
		   }
	    }
    }
}