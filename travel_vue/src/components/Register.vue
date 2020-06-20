<template>
    <div class="container">
        <div class="login_box">
            <el-form class="login_form" :model="registerForm" :rules="registerFormRules" ref="registerFormRef">
                <el-form-item prop="email">
                    <el-input prefix-icon="iconfont el-icon-user" v-model="registerForm.email"/>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input prefix-icon="iconfont el-icon-s-cooperation" v-model="registerForm.password"
                              type="password"/>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input prefix-icon="iconfont el-icon-s-cooperation" v-model="registerForm.rpassword"
                              type="password"/>
                </el-form-item>
                <el-form-item prop="code">
                    <el-input v-model="registerForm.code" prefix-icon="iconfont el-icon-apple" class="code-input"/>
                    <el-button type="success" plain class="code-btn" @click="getCode" :disabled="codeBtnDisable">
                        {{codeBtnText}}
                    </el-button>
                </el-form-item>
                <el-form-item prop="allow">
                    <el-checkbox label="我已阅读并同意相关服务条款和隐私政策" v-model="registerForm.allow"/>
                </el-form-item>
                <el-form-item class="buttons">
                    <el-button type="primary" @click="register">注册</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-row type="flex" justify="space-around">
            <el-col :span="7">
                <router-link :to="{name: 'login'}" class="link">立即登录</router-link>
            </el-col>
            <el-col :span="7">
                <router-link class="link" :to="{name: 'login'}">二维码登录</router-link>
            </el-col>
        </el-row>
    </div>


</template>

<script>


    export default {
        name: "Register",
        data() {
            var checkEmail = (rule, value, callback) => {
                const regEmail = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/;
                if (regEmail.test(value)) {
                    return callback()
                }
                callback(new Error("请输入合法的邮箱"))
            };
            return {
                registerForm: {
                    email: '',
                    password: '',
                    rpassword: '',
                    code: '',
                    allow: '',
                },
                registerFormRules: {
                    email: [
                        {required: true, message: '请输入邮箱', trigger: 'blur'},
                        {validator: checkEmail, trigger: "blur"}
                    ],
                    password: [
                        {required: true, message: '请输入密码', trigger: 'blur'},
                        {min: 6, max: 15, trigger: 'blur', message: '长度在6-15之间'}
                    ],
                    rpassword: [
                        {required: true, message: '请再次输入密码', trigger: 'blur'},
                    ],
                    code: [
                        {required: true, message: '请输入验证码', trigger: 'blur'},
                    ],
                    allow: [
                        {required: true, message: '请同意协议', trigger: 'blur'},
                    ]
                },
                timeDown: 0,
                codeBtnText: "点击获取验证码",
                codeBtnDisable: false,
                timeId: null,
                waitTime: 10,
                getEmailCode: false,
            }
        },
        methods: {
            getCode() {
                // 验证邮箱字段
                this.$refs.registerFormRef.validateField('email', async valid => {
                    if (!valid) {
                        this.getEmailCode = true;
                        // todo 请求发送邮件接口
                        const {data: resp} = await this.$http.post('account/send-mail', {email: this.registerForm.email});
                        if (resp.success) {
                            this.$message({
                                message: resp.message,
                                center: true,
                                type: "success"
                            });
                            // 开启定时器
                            this.timeId = setInterval(this.codeBtnActive, 1000);
                        } else {
                            this.$message({
                                message: resp.message,
                                center: true,
                                type: "error"
                            });
                        }
                    }
                })
            },
            register() {
                this.$refs.registerFormRef.validate(async valid => {
                    if (valid) {

                        // 校验是否获取了验证码
                        if (!this.getEmailCode) {
                            return this.$message({
                                message: "请先获取验证码!",
                                center: true,
                                type: "warning"
                            });
                        }


                        const {data: resp} = await this.$http.post('account/register', this.registerForm);
                        if (resp.success) {
                            this.$message({
                                message: "注册成功!",
                                center: true,
                                type: "success"
                            });
                            await this.$router.push("login")
                        } else {
                            this.$message({
                                message: resp.message,
                                center: true,
                                type: "error"
                            });
                        }
                    }
                })
            },
            codeBtnActive() {
                if (this.timeDown === this.waitTime) {
                    this.codeBtnText = "点击获取验证码";
                    this.codeBtnDisable = false;
                    this.timeDown = 0;
                    // 定时器清除
                    clearInterval(this.timeId)
                } else {
                    this.codeBtnDisable = true;
                    this.timeDown++;
                    this.codeBtnText = `重新发送${this.waitTime - this.timeDown}s`;
                }
            }
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

    .link {
        color: #606266;
        font-size: 14px;
        text-decoration: none;
    }

    .code-input {
        width: 52%;
    }


    .el-button {
        width: 100%;
    }

    .el-row {
        left: 20px;
    }

    .code-btn {
        position: absolute;
        width: 45%;
        right: 0;
    }

</style>
