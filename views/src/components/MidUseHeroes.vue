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
			var param={version:this.$route.params.version,playerid:this.$route.params.playerid}
			this.$store.dispatch('reqMidUseHeroes',param)
		}
	}
</script>
<style>
</style>