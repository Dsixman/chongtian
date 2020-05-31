<template>
	<div>
		
	</div>
</template>
<script>
	export default{
		data(){
			return {

			}
		},
		created(){
			var param={version:this.$route.query.version,playerid:this.$route.query.playerid}
			this.$store.dispatch('getMidUseHeroes',param)
		}
	}
</script>
<style>
</style>