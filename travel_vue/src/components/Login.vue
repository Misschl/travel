<template>
    <div class="container">
        <div class="title">记录你的心路历程</div>
        <div class="login_box">

            <el-form class="login_form" :model="loginForm" :rules="loginFormRules" ref="loginFormRef">
                <el-form-item prop="username">
                    <el-input prefix-icon="iconfont el-icon-user" v-model="loginForm.username"/>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input prefix-icon="iconfont el-icon-s-cooperation" v-model="loginForm.password" type="password"/>
                </el-form-item>
                <el-form-item class="">
                    <el-row type="flex" class="row-bg" justify="space-between">
                        <el-col>
                            <el-checkbox label="记住用户名" name="type"/>
                        </el-col>
                        <el-col>
                            <el-checkbox label="记住密码" name="type"/>
                        </el-col>
                        <el-col>
                            <el-link href="" target="_blank">找回密码</el-link>
                        </el-col>
                    </el-row>
                </el-form-item>
                <el-form-item class="buttons">
                    <el-button type="primary" @click="login">登录</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-row type="flex"  justify="space-around">
            <el-col :span="7">
                <el-link href="" target="_blank">立即注册</el-link>
            </el-col >
            <el-col :span="7">
                <el-link href="" target="_blank">二维码登录</el-link>
            </el-col>
        </el-row>
    </div>


</template>

<script>


    export default {
        name: "Login",
        data() {
            return {
                loginForm: {
                    username: '',
                    password: ''
                },
                loginFormRules: {
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
    .title {
        color: #67C23A;
        text-align: center;
        margin-bottom: 10%;
        font-size: 20px;
    }
</style>
