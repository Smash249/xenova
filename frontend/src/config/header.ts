import {
  Award,
  BookOpen,
  Clock,
  Cpu,
  Factory,
  Globe,
  Handshake,
  Headphones,
  Settings,
  Settings2,
  ShieldCheck,
  Target,
  Users,
  Zap,
} from "lucide-vue-next"

import type { SystemConfig } from "@/types/layout"

export const systemConfig: SystemConfig = {
  navItems: [
    {
      title: "首页",
      path: "/dashboard",
      isNav: true,
    },
    {
      title: "关于我们",
      isNav: false,
    },
    {
      title: "产品中心",
      path: "/product",
      isNav: true,
    },
    { title: "配件商城", path: "/accessory", isNav: true },
    {
      title: "企业动态",
      path: "/news",
      isNav: true,
    },
    {
      title: "下载中心",
      path: "/download",
      isNav: true,
    },
    {
      title: "荣誉和文化",
      path: "/honor",
      isNav: true,
    },
    {
      title: "联系我们",
      path: "/contact",
      isNav: true,
    },
  ],

  products: [
    {
      id: "01",
      name: "视觉激光打标机",
      subtitle: "High-Precision Laser Marking",
      features: [
        "超视觉激光自动校准系统",
        "AI 智能领航，全自动无人工干预",
        "模块化集成，支持流水线无缝对接",
      ],
      image:
        "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?auto=format&fit=crop&q=80&w=800&h=600",
    },
    {
      id: "02",
      name: "全自动点胶机器人",
      subtitle: "Automated Dispensing Robot",
      features: [
        "微米级高精度视觉定位",
        "六轴联动控制，复杂路径规划",
        "实时胶量监控与智能补偿",
      ],
      image:
        "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?auto=format&fit=crop&q=80&w=800&h=600",
    },
    {
      id: "03",
      name: "精密焊接平台",
      subtitle: "Precision Soldering Platform",
      features: [
        "PID 智能恒温焊接技术",
        "在线 AOI 焊点质量检测",
        "柔性化生产配置，快速换型",
      ],
      image:
        "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?auto=format&fit=crop&q=80&w=800&h=600",
    },
  ],

  advantages: [
    {
      id: "01",
      tabTitle: "实力优势",
      icon: ShieldCheck,
      subTitle: "10年行业经验",
      headline: "专业研发，涉足各行业领域自动化设备",
      descriptions: [
        "多年生产制造研发非标自动化设备，拥有1500平方厂房，是集研发制造于一体的行业标杆。",
        "产品种类多，适用于多个行业，是业界的工业自动化系统解决方案，我们以质量作为核心竞争力。",
      ],
      image:
        "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?q=80&w=1000&auto=format&fit=crop",
      stats: [
        { label: "研发厂房", value: "1500㎡", icon: Factory },
        { label: "行业经验", value: "10 Years", icon: Award },
      ],
    },
    {
      id: "02",
      tabTitle: "产品优势",
      icon: Cpu,
      subTitle: "10年行业经验",
      headline: "产品种类多，适用于多种行业",
      descriptions: [
        "产品种类多，适用于多个行业，是业界的工业自动化系统解决方案，我们以质量作为核心竞争力。",
        "产品种类多，适用于多个行业，是业界的工业自动化系统解决方案，我们以质量作为核心竞争力。",
      ],
      image:
        "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?q=80&w=1000&auto=format&fit=crop",
      stats: [
        { label: "覆盖行业", value: "20+", icon: Globe },
        { label: "产品良率", value: "99.9%", icon: ShieldCheck },
      ],
    },
    {
      id: "03",
      tabTitle: "深度定制",
      icon: Settings,
      subTitle: "10年行业经验",
      headline: "依据客户要求定制化需求服务",
      descriptions: [
        "产品种类多，适用于多个行业，是业界的工业自动化系统解决方案，我们以质量作为核心竞争力。",
        "产品种类多，适用于多个行业，是业界的工业自动化系统解决方案，我们以质量作为核心竞争力。",
      ],
      image:
        "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?q=80&w=1000&auto=format&fit=crop",
      stats: [
        { label: "定制方案", value: "1v1", icon: Settings },
        { label: "响应速度", value: "24h", icon: Zap },
      ],
    },
    {
      id: "04",
      tabTitle: "售后服务",
      icon: Headphones,
      subTitle: "10年行业经验",
      headline: "向客户提供全面的技术咨询服务",
      descriptions: [
        "产品种类多，适用于多个行业，是业界的工业自动化系统解决方案，我们以质量作为核心竞争力。",
        "产品种类多，适用于多个行业，是业界的工业自动化系统解决方案，我们以质量作为核心竞争力。",
      ],
      image:
        "https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?q=80&w=1000&auto=format&fit=crop",
      stats: [
        { label: "技术支持", value: "All Day", icon: Headphones },
        { label: "客户满意", value: "100%", icon: ShieldCheck },
      ],
    },
  ],

  chooseUs: {
    merit: [
      {
        id: "01",
        title: "24小时快速响应",
        desc: "全天候技术支持，售后无忧",
        icon: Clock,
        color: "bg-blue-500",
      },
      {
        id: "02",
        title: "高精尖技术团队",
        desc: "硕博领衔，深耕自动化领域",
        icon: Users,
        color: "bg-indigo-500",
      },
      {
        id: "03",
        title: "深度开发高效合理",
        desc: "底层算法优化，拒绝性能冗余",
        icon: Cpu,
        color: "bg-cyan-500",
      },
      {
        id: "04",
        title: "ODM定制化需求",
        desc: "从概念到量产的端到端服务",
        icon: Settings2,
        color: "bg-slate-600",
      },
    ],
    values: [
      { text: "专业知识", icon: BookOpen },
      { text: "精准需求", icon: Target },
      { text: "互利共赢", icon: Handshake },
    ],
  },
  footerItems: [
    "包装机自动化",
    "包装自动化",
    "医疗设备自动化",
    "食品行业自动化",
    "平台点胶机螺丝机",
    "自动化流水线物流线",
    "自研气密测试仪",
    "协作机器人",
    "企业咨询管理",
  ],
}
