<template>
    <div class="container">
        <div class="login_box">
            <el-form class="login_form" :model="registerForm" :rules="registerFormRules" ref="registerFormRef">
                <el-form-item prop="username">
                    <el-input prefix-icon="iconfont el-icon-user" v-model="registerForm.username"/>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input prefix-icon="iconfont el-icon-s-cooperation" v-model="registerForm.password"
                              type="password"/>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input prefix-icon="iconfont el-icon-s-cooperation" v-model="registerForm.rpassword"
                              type="password"/>
                </el-form-item>
                <el-form-item>
                    <el-checkbox label="我已阅读并同意相关服务条款和隐私政策" name="type"/>
                </el-form-item>
                <el-form-item class="buttons">
                    <el-button type="primary" @click="register">注册</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-row type="flex" justify="space-around">
            <el-col :span="7">
                <el-link href="" target="_blank">立即登录</el-link>
            </el-col>
            <el-col :span="7">
                <el-link href="" target="_blank">二维码登录</el-link>
            </el-col>
        </el-row>
    </div>


</template>

<script>


    export default {
        name: "Register",
        data() {
            return {
                registerForm: {
                    username: '',
                    password: '',
                    rpassword: '',
                },
                registerFormRules: {
                    username: [
                        {required: true, message: '请输入用户名', trigger: 'blur'},
                        {min: 3, max: 10, trigger: 'blur', message: '长度在3-10之间'}
                    ],
                    password: [
                        {required: true, message: '请输入密码', trigger: 'blur'},
                        {min: 6, max: 15, trigger: 'blur', message: '长度在6-15之间'}
                    ]
                }
            }
        },
    }
</script>

<style scoped>
    body {
        padding-top: 40px;
        padding-bottom: 40px;
        background-color: #eee;
    }

    .container {
        margin-top: 30%;
    }

    .form-signin {
        max-width: 330px;
        padding: 15px;
        margin: 0 auto;
    }

    .form-signin .form-signin-heading,
    .form-signin .checkbox {
        margin-bottom: 10px;
    }

    .form-signin .checkbox {
        font-weight: normal;
    }

    .form-signin .form-control {
        position: relative;
        height: auto;
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;
        box-sizing: border-box;
        padding: 10px;
        font-size: 16px;
    }

    .form-signin .form-control:focus {
        z-index: 2;
    }

    .form-signin input[type="email"] {
        margin-bottom: -1px;
        border-bottom-right-radius: 0;
        border-bottom-left-radius: 0;
    }

    .form-signin input[type="password"] {
        margin-bottom: 10px;
        border-top-left-radius: 0;
        border-top-right-radius: 0;
    }

    .el-button {
        width: 100%;
    }

    .el-row {
        left: 20px;
    }
</style>
