import SvgIcon from '@/components/SvgIcon/index.vue'

const allGlbComponents = {
  SvgIcon
}

export default {
  install(app: { component: (arg0: string, arg1: any) => void }) {
    Object.keys(allGlbComponents).forEach((name) => {
      app.component(name, allGlbComponents[name])
    })
  }
}
