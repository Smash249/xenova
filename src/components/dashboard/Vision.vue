<template>
  <section
    id="vision"
    class="relative w-full overflow-hidden bg-slate-50 py-24 font-sans"
  >
    <div
      class="pointer-events-none absolute inset-0 z-0 opacity-40"
      style="
        background-image: radial-gradient(#cbd5e1 1px, transparent 1px);
        background-size: 32px 32px;
      "
    ></div>
    <div class="relative z-10 container mx-auto px-4 md:px-6">
      <!-- 顶部大标题：企业愿景 -->
      <div class="mb-20 flex items-center justify-center gap-4">
        <div class="relative hidden h-px w-24 bg-blue-600 md:block">
          <div
            class="absolute top-1/2 right-0 h-2 w-2 -translate-y-1/2 rounded-full bg-blue-600"
          ></div>
        </div>
        <h2
          class="text-center text-3xl font-extrabold text-slate-900 md:text-5xl"
        >
          企业愿景
          <span
            class="mt-2 block text-2xl font-light text-slate-500 md:mt-0 md:inline md:text-4xl"
            >Trusted. 值得信赖</span
          >
        </h2>
        <div class="relative hidden h-px w-24 bg-blue-600 md:block">
          <div
            class="absolute top-1/2 left-0 h-2 w-2 -translate-y-1/2 rounded-full bg-blue-600"
          ></div>
        </div>
      </div>

      <!-- 架构图区域 -->
      <div class="mb-32 grid grid-cols-1 gap-16 lg:grid-cols-2 lg:gap-8">
        <div class="flex flex-col items-center">
          <div class="relative z-10 mb-12">
            <div
              class="clip-hexagon flex h-40 w-48 flex-col items-center justify-center bg-linear-to-br from-blue-200 to-blue-50 p-4 text-center shadow-lg transition-transform duration-300 hover:scale-105"
            >
              <component
                :is="visionLeft.icon"
                class="mb-2 h-8 w-8 text-blue-700"
              />
              <span class="text-lg font-bold text-slate-800">{{
                visionLeft.title
              }}</span>
            </div>
            <!-- 垂直连接线 -->
            <div
              class="absolute top-full left-1/2 h-12 w-px -translate-x-1/2 bg-blue-300"
            ></div>
          </div>

          <div class="relative flex w-full max-w-md justify-around">
            <div
              class="absolute top-0 right-[25%] left-[25%] h-px bg-blue-300"
            ></div>

            <div
              v-for="(item, idx) in visionLeftChildren"
              :key="idx"
              class="group relative flex flex-col items-center pt-8"
            >
              <div
                class="absolute top-0 left-1/2 h-8 w-px -translate-x-1/2 bg-blue-300"
              ></div>

              <div
                class="clip-hexagon flex h-36 w-40 flex-col items-center justify-center border-b-4 border-blue-500 bg-white p-2 text-center shadow-md transition-transform duration-300 group-hover:-translate-y-2"
              >
                <component :is="item.icon" class="mb-1 h-6 w-6 text-blue-500" />
                <span class="font-bold text-slate-800">{{ item.title }}</span>
                <span class="mt-1 text-xs text-slate-500">{{
                  item.subtitle
                }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 右侧树状图 -->
        <div class="flex flex-col items-center">
          <div class="relative z-10 mb-12">
            <div
              class="clip-hexagon flex h-40 w-48 flex-col items-center justify-center bg-linear-to-br from-cyan-200 to-cyan-50 p-4 text-center shadow-lg transition-transform duration-300 hover:scale-105"
            >
              <component
                :is="visionRight.icon"
                class="mb-2 h-8 w-8 text-cyan-700"
              />
              <span class="text-lg font-bold text-slate-800">{{
                visionRight.title
              }}</span>
            </div>

            <div
              class="absolute top-full left-1/2 h-12 w-px -translate-x-1/2 bg-cyan-300"
            ></div>
          </div>

          <div
            class="relative flex w-full flex-wrap justify-center gap-2 md:gap-4"
          >
            <div
              class="absolute top-0 right-[10%] left-[10%] hidden h-px bg-cyan-300 md:block"
            ></div>

            <div
              v-for="(item, idx) in visionRightChildren"
              :key="idx"
              class="group relative flex flex-col items-center pt-0 md:pt-8"
            >
              <div
                class="absolute top-0 left-1/2 hidden h-8 w-px -translate-x-1/2 bg-cyan-300 md:block"
              ></div>

              <div
                class="clip-hexagon flex h-28 w-32 flex-col items-center justify-center border-b-4 border-cyan-500 bg-white p-1 text-center shadow-md transition-transform duration-300 group-hover:-translate-y-2 md:h-28 md:w-32"
              >
                <span class="text-sm font-bold text-slate-800 md:text-base">{{
                  item.title
                }}</span>
                <span
                  class="mt-1 scale-90 text-[10px] text-slate-500 md:text-xs"
                  >{{ item.subtitle }}</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="mb-16 flex items-center justify-center gap-4">
        <div class="relative hidden h-px w-32 bg-blue-600 md:block">
          <div
            class="absolute top-1/2 right-0 h-2 w-2 -translate-y-1/2 rounded-full bg-blue-600"
          ></div>
        </div>
        <h3
          class="px-4 text-center text-2xl font-bold text-slate-800 md:text-3xl"
        >
          赓续初心 探索可持续发展
        </h3>
        <div class="relative hidden h-px w-32 bg-blue-600 md:block">
          <div
            class="absolute top-1/2 left-0 h-2 w-2 -translate-y-1/2 rounded-full bg-blue-600"
          ></div>
        </div>
      </div>

      <div
        class="mx-auto grid max-w-6xl grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 lg:gap-8"
      >
        <div
          v-for="(item, index) in coreValues"
          :key="index"
          class="group flex h-32 overflow-hidden rounded-lg border border-slate-200 bg-white shadow-sm transition-all duration-300 hover:-translate-y-1 hover:shadow-xl md:h-40"
        >
          <div
            class="relative flex w-1/2 flex-col justify-center border-l-4 p-4 md:p-6"
            :class="item.borderColor"
          >
            <h4
              class="mb-2 text-lg font-bold text-slate-800 transition-colors group-hover:text-blue-600 md:text-xl"
            >
              {{ item.title }}
            </h4>
            <p
              class="line-clamp-2 text-xs text-slate-500 md:line-clamp-none md:text-sm"
            >
              {{ item.desc }}
            </p>
          </div>

          <div
            class="relative flex w-1/2 items-center justify-center overflow-hidden"
            :class="item.color"
          >
            <div
              class="absolute -right-4 -bottom-4 h-24 w-24 rounded-full bg-white/20 blur-xl"
            ></div>

            <component
              :is="item.icon"
              class="h-10 w-10 transform text-white transition-all duration-500 group-hover:scale-110 group-hover:rotate-12 md:h-12 md:w-12"
            />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import {
  Award,
  Briefcase,
  Globe,
  Heart,
  Lightbulb,
  ShieldCheck,
  Sprout,
  Target,
  Users,
  Zap,
} from "lucide-vue-next"
import { h } from "vue"

// 定义六边形节点数据结构
interface Node {
  title: string
  subtitle?: string
  icon?: any
}

// 自定义图标
const HandshakeIcon = h(Globe) // 暂时复用 Globe

// 愿景部分数据
const visionLeft: Node = { title: "创国际品牌", icon: Globe }
const visionLeftChildren: Node[] = [
  { title: "技术领先", subtitle: "定义标准", icon: Award },
  { title: "全球信赖", subtitle: "价值伙伴", icon: HandshakeIcon }, // 使用自定义图标
]

const visionRight: Node = { title: "树百年星实", icon: Sprout }
const visionRightChildren: Node[] = [
  { title: "创新为魂", subtitle: "迭代不止", icon: Zap },
  { title: "人才为星", subtitle: "生生不息", icon: Users },
  { title: "质量如山", subtitle: "信誉为脉", icon: ShieldCheck },
  { title: "文化为根", subtitle: "夯实奋进", icon: Target },
  { title: "生态共荣", subtitle: "责任致远", icon: Heart },
]

// 底部色块数据 (模拟图片底部的6个方块)
const coreValues = [
  {
    title: "客户至上",
    desc: "以客户需求为导向，持续创造价值",
    color: "bg-amber-400",
    borderColor: "border-amber-400",
    icon: Users,
  },
  {
    title: "激情奋斗",
    desc: "永葆创业激情，敢于挑战不可能",
    color: "bg-red-700",
    borderColor: "border-red-700",
    icon: Zap,
  },
  {
    title: "拥抱变化",
    desc: "敏锐洞察趋势，主动寻求突破",
    color: "bg-lime-500",
    borderColor: "border-lime-500",
    icon: Lightbulb,
  },
  {
    title: "诚信正直",
    desc: "坚守商业道德，坦荡做人做事",
    color: "bg-blue-900",
    borderColor: "border-blue-900",
    icon: ShieldCheck,
  },
  {
    title: "精益求精",
    desc: "追求卓越品质，关注每一个细节",
    color: "bg-emerald-300",
    borderColor: "border-emerald-300",
    icon: Award,
  },
  {
    title: "团队合作",
    desc: "共享共担，协同作战达成目标",
    color: "bg-blue-600",
    borderColor: "border-blue-600",
    icon: Briefcase,
  },
]
</script>

<style scoped>
/* 自定义六边形裁剪路径 */
.clip-hexagon {
  clip-path: polygon(50% 0%, 100% 25%, 100% 75%, 50% 100%, 0% 75%, 0% 25%);
}

/* 确保子节点在移动端也能对齐 */
@media (max-width: 768px) {
  .clip-hexagon {
    /* 在移动端稍微调整形状或大小以适应 */
    clip-path: polygon(
      10% 0,
      100% 0,
      100% 100%,
      10% 100%,
      0% 50%
    ); /* 变成箭头形状或者圆角矩形可能更好看，这里保持六边形 */
    clip-path: polygon(50% 0%, 100% 25%, 100% 75%, 50% 100%, 0% 75%, 0% 25%);
  }
}
</style>
