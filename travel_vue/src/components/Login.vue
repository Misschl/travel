<template>
    <div class="container">
        <div class="title">记录你的心路历程</div>
        <div class="login_box">
            <el-form class="login_form" :model="loginForm" :rules="loginFormRules" ref="loginFormRef">
                <el-form-item prop="email">
                    <el-input prefix-icon="iconfont el-icon-user" v-model="loginForm.email"/>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input prefix-icon="iconfont el-icon-s-cooperation" v-model="loginForm.password"
                              type="password"/>
                </el-form-item>
                <el-form-item class="">
                    <el-row type="flex" class="row-bg" justify="space-between">
                        <el-col>
                            <el-checkbox label="记住用户名" v-model="rememberEmail"/>
                        </el-col>
                        <el-col>
                            <el-checkbox label="记住密码" v-model="rememberPassword"/>
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
        <el-row type="flex" justify="space-around">
            <el-col :span="7">
                <!--                <el-link href="" >立即注册</el-link>-->
                <router-link :to="{name:'register'}" class="link">立即注册</router-link>
            </el-col>
            <el-col :span="7">
                <router-link :to="{name:'register'}" class="link">二维码登录</router-link>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    export default {
        name: "Login",
        data() {
            var checkEmail = (rule, value, callback) => {
                const regEmail = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/;
                if (regEmail.test(value)) {
                    return callback()
                }
                callback(new Error("请输入合法的邮箱"))
            };
            return {
                loginForm: {
                    email: '',
                    password: '',
                },
                rememberEmail: false,
                rememberPassword: false,
                loginFormRules: {
                    email: [
                        {required: true, message: '请输入邮箱', trigger: 'blur'},
                        {validator: checkEmail, trigger: "blur"}
                    ],
                    password: [
                        {required: true, message: '请输入密码', trigger: 'blur'},
                        {min: 6, max: 18, trigger: 'blur', message: '长度在6-15之间'}
                    ]
                }
            }
        },
        methods: {
            login() {
                this.$refs.loginFormRef.validate(async valid => {
                    if (valid) {
                        const {data: resp} = await this.$http.post('account/login', this.loginForm);
                        if (resp.success) {
                            this.remember();
                            window.sessionStorage.setItem("token", resp.result);
                            this.$message({
                                message: resp.message,
                                center: true,
                                type: "success"
                            });
                            await this.$router.push('index')
                        } else {
                            this.$message({
                                message: resp.message,
                                center: true,
                                type: "error"
                            });
                        }
                    }

                });

            },
            remember() {
                const email = this.rememberEmail ? this.loginForm.email : "";
                const password = this.rememberPassword ? this.loginForm.password : "";
                window.localStorage.setItem("email", email);
                window.localStorage.setItem("password", password);
            }
        },
        created() {
            this.loginForm.email = window.localStorage.getItem("email");
            this.loginForm.password = window.localStorage.getItem("password");
            this.rememberEmail = !!this.loginForm.email;
            this.rememberPassword = !!this.loginForm.password;
        }
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

    .link {
        color: #606266;
        font-size: 14px;
        text-decoration: none;
    }
</style>
