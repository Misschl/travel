<template>
    <div class="user-container">
        <div class="head"></div>
        <div class="h-box">
            <el-card class="h-h-box">
                <el-row>
                    <!--头像区域-->
                    <el-col :span="8">
                        <el-image
                                style="width: 100px; height: 100px"
                                :src="userInfo.avatar"
                                fit="fit"
                                :preview-src-list="[userInfo.avatar]"
                                class="avatar">
                        </el-image>
                    </el-col>
                    <el-col :span="14">
                        <el-col style="font-size: 40px">{{userInfo.username}}</el-col>
                        <el-col style="bottom: 0">
                            <hr>
                        </el-col>
                        <el-col>
                            {{userInfo.sign}}
                        </el-col>
                    </el-col>

                </el-row>


            </el-card>
        </div>
        <div class="j-box">
            <el-row :gutter="25" type="flex" class="row-bg" justify="center" style="padding-top: 20px">
                <el-col :span="6" v-for="item in relateNote" :key="item.id">
                    <div class="menu" style="text-align: center">
                        <div>{{item.count}}</div>
                        <span>{{item.name}}</span>
                    </div>
                </el-col>
            </el-row>
        </div>
        <div class="d-box">
            <el-row justify="space-between" type="flex" class="row-bg" :gutter="8">
                <el-col :span="16">
                    <el-card style="height: 80px;">运动数据</el-card>
                </el-col>
                <el-col :span="16">
                    <el-card style="height: 80px;">健康数据</el-card>
                </el-col>
            </el-row>
            <el-card style="height: 200px; margin-top: 20px">
                地图模块
            </el-card>
        </div>
    </div>
</template>

<script>
    export default {
        name: "User",
        data() {
            return {
                userInfo: {
                    avatar: "https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg",
                    username: "万总",
                    sign: "我是万总"
                },
                relateNote: [
                    {name: '关注', count: 6, id: 1},
                    {name: '粉丝', count: 10, id: 2},
                    {name: '动态', count: 13, id: 3},
                ]
            }
        }
    }
</script>

<style scoped>

    .user-container {
        height: 100%;
    }

    .head {
        height: 5%;
    }

    .h-box {
        height: 25%;
        /*background-color: green;*/
    }

    .j-box {
        height: 15%;
        /*background-color: yellow;*/
    }

    .d-box {
        height: 55%;
        /*background-color: red;*/
    }

    .h-h-box {
        height: 92%;
        width: 100%;
        margin: auto;
    }

    .j-box .el-row .el-col {
        display: flex;
        align-items: center;
    }

</style>
