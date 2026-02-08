import type { LucideIcon } from "lucide-vue-next"

interface SystemConfig {
  navItems: HeaderNavItem[]
  products: ProductItem[]
  footerItems: string[]
  advantages: Advantages[]
  chooseUs: {
    merit: ChooseUsMeritItem[]
    values: ChooseUsValueItem[]
  }
}

interface HeaderNavItem {
  title: string
  path?: string
  isNav: boolean
}

interface ProductItem {
  id: string
  name: string
  subtitle: string
  features: string[]
  image: string
}

interface AdvantagesStatItem {
  label: string
  value: string
  icon: LucideIcon
}

interface Advantages {
  id: string
  tabTitle: string
  icon: LucideIcon
  subTitle: string
  headline: string
  descriptions: string[]
  image: string
  stats: AdvantagesStatItem[]
}

interface ChooseUsMeritItem {
  id: string
  title: string
  desc: string
  icon: LucideIcon
  color: string
}

interface ChooseUsValueItem {
  text: string
  icon: LucideIcon
}

export type { SystemConfig, HeaderNavItem }
