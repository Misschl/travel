<template>
    <div class="container">
        <div class="user-info">
            <span class="avatar">
                <el-avatar size="large"
                           src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"/>
            </span>
            <span>kid123456</span>
        </div>
    </div>
</template>

<script>
    export default {
        name: "User"
    }
</script>

<style scoped>

    .container {
        width: 100%;
        height: 100%;
    }

    .user-info {
        margin-top: 10%;
        width: 92%;
        margin-left: auto;
        height: 22%;
    }
</style>
