import SvgIcon from './global/SvgIcon.vue'

const allGlbComponents: { [key: string]: ComponentOptions } = {
  SvgIcon
}
// export default {
//   install(app) {
//     Object.keys(allGlbComponents).forEach((name) => {
//       app.component(name, allGlbComponents[name])
//     })
//   }
// }

export default {
  install(app: { component: (arg0: string, arg1: any) => void }) {
    Object.keys(allGlbComponents).forEach((name) => {
      app.component(name, allGlbComponents[name])
    })
  }
}