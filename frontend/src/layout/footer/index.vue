<template>
  <footer
    class="border-t border-slate-200 bg-slate-50 pt-16 font-sans text-slate-700"
  >
    <div class="container mx-auto px-6 lg:px-8">
      <div class="grid grid-cols-1 gap-12 lg:grid-cols-12 lg:gap-8">
        <div class="space-y-6 lg:col-span-5">
          <h2 class="text-2xl font-bold text-slate-900">星实科技</h2>
          <div
            class="space-y-4 text-justify text-sm leading-relaxed text-slate-600"
          >
            <p>
              星实（厦门）科技有限公司，是一家坐落于厦门集美软件园三期的高新技术企业。自成立以来，我们始终深耕于非标自动化领域，专注于研发与制造智能机器人、高精度平台（锁付、焊锡、点胶、灌胶）设备及一体化产线，致力于为制造业客户提供高效、可靠的自动化生产解决方案。
            </p>
            <p>
              凭借一支充满活力的年轻研发团队与超过1200平方米的研发制造基地，我们将国际先进技术与国内实际工艺需求深度融合，形成了自主研发与快速交付的核心能力。从专业研发、精益制造到销售与全周期售后服务，星实科技已构建完整的价值链，旨在切实帮助客户解决生产难题、提升效率、降低成本，助力企业实现智能化升级。
            </p>
            <p class="font-medium text-slate-800">
              我们的使命是：以创新的自动化技术，赋能制造，助力客户成功。
            </p>
            <p>欢迎垂询，期待与您共赴智造未来。</p>
          </div>

          <div class="flex gap-4 pt-4">
            <SocialIcon color="bg-green-500">
              <svg class="h-6 w-6 fill-current" viewBox="0 0 24 24">
                <path
                  d="M8 17v-4h-4v-4h4v-2h-6v8h6zm10-6v-6h-6v6h6zm-2-4h-2v2h2v-2zm-12 12h18v-16h-18v16z"
                />
              </svg>
            </SocialIcon>
            <SocialIcon color="bg-black">
              <svg class="h-5 w-5 fill-current" viewBox="0 0 24 24">
                <path
                  d="M12.5 0v18c0 3.3-2.7 6-6 6s-6-2.7-6-6c0-3.3 2.7-6 6-6 .5 0 1 .1 1.5.2v5.2c-.2-.1-.4-.2-.5-.2-1.7 0-3 1.3-3 3s1.3 3 3 3 3-1.3 3-3v-12.5h5v4.5h-2.5z"
                />
              </svg>
            </SocialIcon>
            <SocialIcon color="bg-red-500">
              <svg class="h-6 w-6 fill-current" viewBox="0 0 24 24">
                <path
                  d="M12 2c5.5 0 10 4.5 10 10s-4.5 10-10 10-10-4.5-10-10 4.5-10 10-10zm0 18c4.4 0 8-3.6 8-8s-3.6-8-8-8-8 3.6-8 8 3.6 8 8 8z"
                />
              </svg>
            </SocialIcon>
            <SocialIcon color="bg-blue-600">
              <Mail class="h-5 w-5" />
            </SocialIcon>
          </div>
        </div>

        <div class="lg:col-span-3 lg:pl-8">
          <h3 class="mb-6 text-xl font-bold text-slate-900">产品中心</h3>
          <ul class="space-y-3">
            <li
              v-for="(item, index) in systemConfig.footerItems"
              :key="index"
              class="group flex cursor-pointer items-center"
            >
              <span
                class="mr-3 h-2 w-2 rounded-sm bg-blue-600 transition-all duration-300 group-hover:w-3"
              ></span>
              <span
                class="text-slate-600 transition-colors hover:text-blue-700"
                >{{ item }}</span
              >
            </li>
          </ul>
        </div>
        <div class="lg:col-span-4">
          <h3 class="mb-6 text-xl font-bold text-slate-900">联系我们</h3>
          <ul class="space-y-4 text-slate-600">
            <li class="flex items-start">
              <MapPin class="mt-0.5 mr-3 h-5 w-5 shrink-0 text-blue-600" />
              <span>福建省厦门市集美区软件园三期F8-1888号</span>
            </li>
            <li class="flex items-center">
              <Phone class="mr-3 h-5 w-5 shrink-0 text-blue-600" />
              <span>+86 0592-6818878</span>
            </li>
            <li class="flex items-center">
              <Printer class="mr-3 h-5 w-5 shrink-0 text-blue-600" />
              <span>+86 0592-6818878 (传真)</span>
            </li>
            <li class="flex items-center">
              <Mail class="mr-3 h-5 w-5 shrink-0 text-blue-600" />
              <a
                href="mailto:info@xingshi-tech.com"
                class="transition-colors hover:text-blue-700"
                >info@xingshi-tech.com</a
              >
            </li>
            <li class="flex items-center">
              <Clock class="mr-3 h-5 w-5 shrink-0 text-blue-600" />
              <span>周一至周五 9:00-18:00</span>
            </li>
          </ul>

          <div class="mt-8 border-t border-slate-200 pt-6">
            <p class="mb-4 text-center font-medium text-slate-900">
              欢迎扫码联系我们
            </p>
            <div class="flex justify-center gap-6">
              <div class="text-center">
                <div
                  class="h-28 w-28 rounded-lg border border-slate-100 bg-white p-2 shadow-sm"
                >
                  <img
                    src="https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=WeChat"
                    alt="微信二维码"
                    class="h-full w-full object-contain opacity-80"
                  />
                </div>
                <span class="mt-2 block text-xs text-slate-500"
                  >微信公众号</span
                >
              </div>
              <div class="text-center">
                <div
                  class="h-28 w-28 rounded-lg border border-slate-100 bg-white p-2 shadow-sm"
                >
                  <img
                    src="https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=Contact"
                    alt="联系二维码"
                    class="h-full w-full object-contain opacity-80"
                  />
                </div>
                <span class="mt-2 block text-xs text-slate-500">联系客服</span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="mt-16 border-t border-slate-200 py-8">
        <div
          class="flex flex-col items-center justify-between gap-4 text-xs text-slate-500 md:flex-row"
        >
          <div class="space-y-1 text-center md:text-left">
            <p>
              © 2021-2026 星实(厦门)科技有限公司 版权所有 | 闽ICP备688888889号
            </p>
            <p>
              电子营业执照 | 增值电信业务经营许可证：B1.B2-20200593 | 闽公网安备
              520888888号
            </p>
          </div>

          <div class="flex gap-6">
            <a href="#" class="transition-colors hover:text-blue-600"
              >法律申明</a
            >
            <a href="#" class="transition-colors hover:text-blue-600"
              >隐私政策申明</a
            >
            <a href="#" class="transition-colors hover:text-blue-600"
              >Cookies</a
            >
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { Clock, Mail, MapPin, Phone, Printer } from "lucide-vue-next"

import { systemConfig } from "@/config/header"

const SocialIcon = {
  props: {
    color: String,
  },
  template: `
    <div :class="['w-10 h-10 rounded-full flex items-center justify-center text-white transition-transform hover:scale-110 cursor-pointer', color]">
      <slot></slot>
    </div>
  `,
}
</script>
