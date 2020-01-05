<template>
    <div class="formwarpper">
        
        <el-form ref="form" :model="form" label-width="80px" :rules="rules">
          <!--   <el-form-item label="用户名:" :rules="{required: true, message: '用户名不能为空', trigger: 'blur'}" >
          <el-input v-model="form.username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码:" :rules="{required: true, message: '密码不能为空', trigger: 'blur'}">
          <el-input v-model="form.password"></el-input>
                    </el-form-item>
                    <el-form-item label="验证码:" :rules="{required: true, message: '验证码不能为空', trigger: 'blur'}">
          <el-input v-model="form.captcha"></el-input>
                    </el-form-item> -->
          <el-form-item label="">
            <el-input v-model="form.captchakey" type="hidden" value=""></el-input>
          </el-form-item>
          <el-form-item label="用户名:">
            <el-input v-model="form.username" type="text"></el-input>
          </el-form-item>
          <el-form-item label="密码:">
            <el-input v-model="form.password" type="password"></el-input>
          </el-form-item>
          <el-form-item label="验证码:">
            <el-input v-model="form.captcha" type="text"></el-input>
          </el-form-item>
          <p><img :src="this.$store.state.captchaStr" @click="getCaptcha"></p>
        <el-form-item>
            <el-button class="login-button" type="primary" @click="submitForm('form')">登录</el-button>
        </el-form-item>
        </el-form>
    	<!-- <form action="./api/login">
            <input type="hidden" name="captchakey" v-bind:value="captchaKey">
                    用户名：<input type="text" name="username" value="">
                    密码：<input type="text" name="password" value="">
                    验证码：<input type="text" name="captcha"><img :src="captchaStr" @click="getCaptcha">
        </form> -->
    </div>
</template>

<script type="ecmascript-6">
import Vue from 'vue'
/*import {mapState} from 'vuex'*/
    export default{
        data(){
            return{
                form:{
                    username:'',
                    password:'',
                    captcha:'',
                    captchakey:''
                },
                rules:{
                  username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                  ],
                  password: [
                    { required: true, message: '请输入密码', trigger: 'blur' }
                  ],
                  captcha: [
                    { required: true, message: '请输入验证码', trigger: 'blur' }
                  ],
                }
            }
        },
    	computed:{
    		/*...mapState(['loginState']),*/
    	},
        created(){
            this.$store.dispatch('getLoginData')
        },
        methods:{
            getCaptcha(){
                this.$store.dispatch('getCaptcha')
            },
            submitForm(formName) {
                this.form.captchakey=this.$store.state.captchaKey
                this.$refs[formName].validate((valid) => {
                    if (valid) { 
                        Vue.axios.post("./api/login",JSON.stringify(this.form))
                        .then((data)=>{
                            var result={}
                             result.loginState=data.data.loginState
                             result.userName=data.data.userName
                            this.$store.dispatch('getLoginAuth',result)
                            if (result.loginState==true){
                              this.$router.push('/hero')
                            }
                        }).catch((err)=>{
                            console.log(err)
                        })
                    }else{
                        console.log('error submit!!');
                        return false;
                    }
                })//validate
            }
        }
    
    }
</script>

<style rel="stylesheet">
    .formwarpper{
        width:1050px;margin:0 auto 0 auto;
    }
</style>