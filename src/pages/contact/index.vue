<template>
  <div
    class="relative flex min-h-screen w-full flex-col bg-slate-50 font-sans text-slate-600"
  >
    <div
      class="relative flex h-[40vh] items-center justify-center overflow-hidden bg-slate-900"
    >
      <!-- 背景装饰 -->
      <div
        class="absolute inset-0 z-10 bg-linear-to-r from-blue-900/80 to-slate-900/80"
      ></div>

      <div
        class="animate-pulse-slow absolute inset-0 bg-[url('./banner/about.webp')] bg-cover bg-center"
      ></div>

      <div class="relative z-20 px-4 text-center">
        <h1
          class="mb-4 text-4xl font-bold tracking-wider text-white drop-shadow-md md:text-5xl"
        >
          联系我们
        </h1>
      </div>
    </div>

    <!-- 主要内容区域 -->
    <div class="container mx-auto px-4 py-12 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 gap-12 lg:grid-cols-12 lg:gap-12">
        <!-- 第一列：公司简介 -->
        <div class="lg:col-span-5">
          <h2 class="mb-8 text-3xl font-black tracking-tight text-slate-900">
            <span
              class="bg-linear-to-r from-blue-700 to-indigo-600 bg-clip-text text-transparent"
              >星实科技</span
            >
          </h2>
          <div class="space-y-6 text-justify leading-relaxed text-slate-600">
            <p>
              星实（厦门）科技有限公司，是一家坐落于厦门集美软件园三期的高新技术企业。自成立以来，我们始终深耕于非标自动化领域，专注于研发与制造智能机器人、高精度平台（锁付、焊锡、点胶、灌胶）设备及一体化产线，致力于为制造业客户提供高效、可靠的自动化生产解决方案。
            </p>
            <p>
              凭借一支充满活力的年轻研发团队与超过1200平方米的研发制造基地，我们将国际先进技术与国内实际工艺需求深度融合，形成了自主研发与快速交付的核心能力。从专业研发、精益制造到销售与全周期售后服务，星实科技已构建完整的价值链，旨在切实帮助客户解决生产难题、提升效率、降低成本，助力企业实现智能化升级。
            </p>
            <div class="rounded-lg border-l-4 border-blue-600 bg-blue-50 p-4">
              <p class="font-bold text-blue-800">
                我们的使命是：以创新的自动化技术，赋能制造，助力客户成功。
              </p>
            </div>
            <p class="font-medium text-slate-900">
              欢迎垂询，期待与您共赴智造未来。
            </p>
          </div>
        </div>

        <!-- 第二列：产品中心 -->
        <div class="lg:col-span-3 lg:border-l lg:border-slate-200 lg:pl-12">
          <h3
            class="mb-8 flex items-center gap-2 text-xl font-bold text-slate-900"
          >
            <div class="h-6 w-1 rounded-full bg-blue-600"></div>
            产品中心
          </h3>
          <ul class="space-y-4">
            <li
              v-for="(item, index) in products"
              :key="index"
              class="group flex items-center gap-3"
            >
              <div
                class="flex h-1.5 w-1.5 items-center justify-center rounded-full bg-slate-300 transition-all duration-300 group-hover:w-3 group-hover:bg-blue-600"
              ></div>
              <a
                href="#"
                class="text-base transition-colors hover:font-bold hover:text-blue-700"
              >
                {{ item }}
              </a>
            </li>
          </ul>
        </div>

        <!-- 第三列：联系我们 -->
        <div class="lg:col-span-4 lg:pl-4">
          <h3
            class="mb-8 flex items-center gap-2 text-xl font-bold text-slate-900"
          >
            <div class="h-6 w-1 rounded-full bg-blue-600"></div>
            联系我们
          </h3>

          <div class="mb-10 space-y-5">
            <div
              class="flex items-start gap-4 rounded-lg bg-slate-50 p-3 transition-colors hover:bg-white hover:shadow-sm"
            >
              <div
                class="mt-1 flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-blue-100 text-blue-600"
              >
                <MapPinIcon class="h-4 w-4" />
              </div>
              <div>
                <div class="text-xs font-bold text-slate-400 uppercase">
                  地址
                </div>
                <div class="font-medium text-slate-800">
                  福建省厦门市集美区软件园三期F8-1888号
                </div>
              </div>
            </div>

            <div
              class="flex items-center gap-4 rounded-lg bg-slate-50 p-3 transition-colors hover:bg-white hover:shadow-sm"
            >
              <div
                class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-blue-100 text-blue-600"
              >
                <PhoneIcon class="h-4 w-4" />
              </div>
              <div>
                <div class="text-xs font-bold text-slate-400 uppercase">
                  电话
                </div>
                <div class="font-mono font-bold text-slate-800">
                  +86 0592-6818878
                </div>
              </div>
            </div>

            <div
              class="flex items-center gap-4 rounded-lg bg-slate-50 p-3 transition-colors hover:bg-white hover:shadow-sm"
            >
              <div
                class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-blue-100 text-blue-600"
              >
                <PrinterIcon class="h-4 w-4" />
              </div>
              <div>
                <div class="text-xs font-bold text-slate-400 uppercase">
                  传真
                </div>
                <div class="font-mono text-slate-800">+86 0592-6818878</div>
              </div>
            </div>

            <div
              class="flex items-center gap-4 rounded-lg bg-slate-50 p-3 transition-colors hover:bg-white hover:shadow-sm"
            >
              <div
                class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-blue-100 text-blue-600"
              >
                <MailIcon class="h-4 w-4" />
              </div>
              <div>
                <div class="text-xs font-bold text-slate-400 uppercase">
                  邮箱
                </div>
                <a
                  href="mailto:info@xingshi-tech.com"
                  class="font-bold text-slate-800 transition-colors hover:text-blue-600"
                >
                  info@xingshi-tech.com
                </a>
              </div>
            </div>

            <div
              class="flex items-center gap-4 rounded-lg bg-slate-50 p-3 transition-colors hover:bg-white hover:shadow-sm"
            >
              <div
                class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-blue-100 text-blue-600"
              >
                <ClockIcon class="h-4 w-4" />
              </div>
              <div>
                <div class="text-xs font-bold text-slate-400 uppercase">
                  工作时间
                </div>
                <div class="text-slate-800">周一至周五 9:00-18:00</div>
              </div>
            </div>
          </div>

          <!-- 二维码区域 -->
          <div class="border-t border-slate-200 pt-8">
            <p
              class="mb-6 text-center text-sm font-bold tracking-wider text-slate-400"
            >
              欢迎扫码联系我们
            </p>
            <div class="flex justify-center gap-8">
              <div class="group flex flex-col items-center gap-3">
                <div
                  class="rounded-xl border border-slate-200 bg-white p-2 shadow-sm transition-all duration-300 group-hover:-translate-y-2 group-hover:border-blue-200 group-hover:shadow-lg group-hover:shadow-blue-100"
                >
                  <img
                    src="https://via.placeholder.com/120x120?text=QR"
                    alt="微信公众号"
                    class="h-28 w-28 object-contain"
                  />
                </div>
                <span
                  class="text-xs font-medium text-slate-500 group-hover:text-blue-600"
                  >微信公众号</span
                >
              </div>

              <div class="group flex flex-col items-center gap-3">
                <div
                  class="rounded-xl border border-slate-200 bg-white p-2 shadow-sm transition-all duration-300 group-hover:-translate-y-2 group-hover:border-blue-200 group-hover:shadow-lg group-hover:shadow-blue-100"
                >
                  <img
                    src="https://via.placeholder.com/120x120?text=QR"
                    alt="联系客服"
                    class="h-28 w-28 object-contain"
                  />
                </div>
                <span
                  class="text-xs font-medium text-slate-500 group-hover:text-blue-600"
                  >联系客服</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  ClockIcon,
  MailIcon,
  MapPinIcon,
  PhoneIcon,
  PrinterIcon,
} from "lucide-vue-next"

const products = [
  "包装机自动化",
  "包装自动化",
  "医疗设备自动化",
  "食品行业自动化",
  "平台点胶机螺丝机",
  "自动化流水线物流线",
  "自研气密测试仪",
  "协作机器人",
  "企业咨询管理",
]
</script>
