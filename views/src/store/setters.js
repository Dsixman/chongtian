export default{
	/*getSecondCat(state){
		var catl=state.catList;
		var seccatl=[];
		for(var key in catl){	
			seccatl.push(catl[key].children)	
		}
		return seccatl;
	}*/
	/*getHeroData(state){
		var heroData=state.heroData
		for (var key in heroData){
			var str=heroData[key]
			if (typeof(str)=="string"){
				var result=str.indexOf(".jpg")
				if (result!=-1){
					var heroIcon=require(""+str)//这样写不行
					heroData[key]=heroIcon	
				}
			}
		}
		return heroData
	}*/
}