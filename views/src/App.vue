<template>
  <el-container id="app">
    <el-container>
      <el-header style="padding:0px">
        <el-menu  class="el-menu-demo" mode="horizontal" style="background:#091723;border:0px;">
          <el-menu-item @click="drawer = true" ><i class="el-icon-s-unfold" style="color:#fff" /></el-menu-item>
          <el-menu-item style="color:#fff"><router-link to="/">战队数据</router-link></el-menu-item>
          <el-menu-item style="color:#fff"><router-link to="/bp">bp模拟</router-link></el-menu-item>
          <el-menu-item ><router-link to="/hero">英雄模拟</router-link> </el-menu-item>
          <div style="float:right;line-height:55px;color:#909399"><i class="el-icon-user"></i>用户：{{this.$store.state.userName}}</div>
            <div style="float:right;margin-top:10px;margin-right:30px">
              <el-input placeholder="请输入内容" v-model="input3" class="input-with-select" style="background:#091723">
                <el-select v-model="select" slot="prepend" placeholder="选择搜索内容">
                  <el-option label="比赛" value="1"></el-option>
                  <el-option label="战队" value="2"></el-option>
                  <el-option label="英雄对线" value="3"></el-option>
                </el-select>
                <el-button slot="append" icon="el-icon-search"></el-button>
              </el-input>
            </div>
          
          <!-- <p>{{this.$store.state.userName}}</p> -->
        </el-menu>
      </el-header>
      <el-container>
      <el-drawer :visible.sync="drawer"  :with-header="false"  direction="ltr">
        <span>我来啦!</span>
      </el-drawer>
        <router-view/>
      </el-container>
    <!--   <el-footer>Footer</el-footer> -->
    </el-container>
</el-container>
</template>
<script>
//import {mapActions} from 'vuex'
import {mapState} from 'vuex'
  export default{
    data() {
      return {
        drawer: false,
        input3: '',
        select: ''
      };
    },
    computed:{
       //...mapState(['userName'])
     },
    created(){
     
     /* var router=this.$router
      this.$store.dispatch('getAuth',router)*/
    },
    methods: {
      /*isLogin(){
        if (!this.$store.state.loginState) {
          this.$route.push("/login")
        }
      }*/
    }
  }
</script>
<style>
a{
  text-decoration:none;
}
body{
  background:#002441

}
.margin-left-200{
  margin-left:-200px
}
.pd-0{
  padding:0px
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #fff;
}
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
 .el-select .el-input {
    width: 130px;
  }
  .input-with-select .el-input-group__prepend {
    background-color: #fff;
  }
</style>
