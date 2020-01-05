<template>
	<el-container>
		<div class="charts" style="width:100%;height:600px;">
			<div id="physical_resistance" style="width:100%;height:600px;">	
			</div>
		</div>
	</el-container>
</template>
<script>
var echarts = require('echarts/lib/echarts');
require('echarts/lib/chart/line');
require('echarts/lib/component/title');
	export default{
		data(){
			return {

			}
		},
		mounted(){
			this.drawPhyResisLine();
		},
		methods:{
			drawPhyResisLine(){
				let lineGraph = echarts.init(document.getElementById('physical_resistance'));
				var xaxisData=[];
				var seriesData=[];
				for (var i=0;i<=50;i++){
					var phyResisPect=0.052*i/(0.9+0.048*i)//物理减伤公式
					phyResisPect=phyResisPect.toFixed(2)
					seriesData.push(phyResisPect)	
					xaxisData.push(i.toString())
				}
				lineGraph.setOption({
					textStyle: {
						color:"#fff"
					},
					title: {
            			text: '物理抗性与护甲',
            			left: 'center',
            			textStyle: {
							color:"#fff"
						}
        			},
					xAxis: {
		        		type: 'category',
		        		data: xaxisData,
		        		boundaryGap: false,
		        		axisLine:{
         					lineStyle:{
            				color:'#fff' //更改坐标轴颜色
         					}
   						}
		    		},
		   			yAxis: {
		        		type: 'value',
		        		
		        		axisLine:{
         					lineStyle:{
            				color:'#fff' //更改坐标轴颜色
         					}
   						}
		    		},
		    		series: [{
		        		data: seriesData,
		        		type: 'line',
		        		smooth: true,
		        		label: {
		                	normal: {
		                    	show: true,
		                    	position: 'top'
		                	}
		            	}
		    		}]
				})
			},
			drawMgiResisLine(){
				  
			}
		}
	}
</script>
<style></style>