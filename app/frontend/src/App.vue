<script setup>
import HelloWorld from './components/HelloWorld.vue'</script>

<template>
  <div id="app">
    <!-- 顶部标题栏 -->
    <header class="header">
      <img src="./assets/images/logo-universal.png" class="logo" alt="Logo" />
      <h1 class="title">爱心平台接力站</h1>
    </header>

    <div class="main-container">
      <!-- 左边栏 -->
      <aside class="sidebar">
        <nav>
          <ul>
            <li>
              <a href="#">迷失列表</a>
              <ul class="sub-menu">
                <li><a href="#">迷失列表 1-1</a></li>
                <li><a href="#">迷失列表 1-2</a></li>
                <li><a href="#">迷失列表 1-3</a></li>
              </ul>
            </li>
            <li>
              <a href="#">最新游戏</a>
              <ul class="sub-menu">
                <li><a href="#">最新游戏 1-1</a></li>
                <li><a href="#">最新游戏 1-2</a></li>
                <li><a href="#">最新游戏 1-3</a></li>
              </ul>
            </li>
            <li>
              <a href="#">系统配置</a>
              <ul class="sub-menu">
                <li><a href="#">系统配置 1-1</a></li>
                <li><a href="#">系统配置 1-2</a></li>
                <li><a href="#">系统配置 1-3</a></li>
              </ul>
            </li>
          </ul>
        </nav>
      </aside>

      <!-- 内容区 -->
      <section class="content">
        <h2>内容区</h2>
        <p>这里是主要内容区域。</p>
        <hello-world ></hello-world>
      </section>
    </div>

    <!-- 底部栏 -->
    <footer class="footer">
      <p>© 2024 版权信息</p>
    </footer>
  </div>
</template>

<style src="./style.css"></style>
