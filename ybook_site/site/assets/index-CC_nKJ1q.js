import{_ as m}from"./_plugin-vue_export-helper-DlAUqK2U.js";import{r as v,o as l,c as $,w as h,p as V,a as I,b as a,d as U,e as _,u as F,f as u,g as t,t as T,h as O,E as z,i as B,F as W,j as Y,k,n as j,l as g,R as G,m as A,q as D,v as J,s as N,x as K,y as Q}from"./index-DMdI58a_.js";import{u as E,g as X}from"./user-2AIJstJP.js";const Z="/ybook/assets/logo-C_FfGPO6.svg",aa={},sa=s=>(V("data-v-ef4ffc62"),s=s(),I(),s),ta=sa(()=>a("img",{src:Z,alt:"logo"},null,-1));function ea(s,e){const d=v("router-link");return l(),$(d,{class:"logo_container",to:"/home"},{default:h(()=>[ta]),_:1})}const ca=m(aa,[["render",ea],["__scopeId","data-v-ef4ffc62"]]),C=s=>(V("data-v-a6cd265a"),s=s(),I(),s),da={class:"login-container"},na={class:"content"},ia=C(()=>a("div",{class:"left"},[a("div",{class:"header"},[a("div",{class:"login-reason"},"登录后推荐更懂你的笔记"),a("img",{src:"",alt:""})]),a("div",{class:"code-area"})],-1)),la={class:"right"},oa=C(()=>a("div",{class:"title"},"邮箱验证码登录",-1)),va={class:"input-container"},ha=["disabled"],_a=C(()=>a("div",{class:"err-msg"},null,-1)),pa=C(()=>a("div",{class:"agree"},null,-1)),ra=C(()=>a("div",{class:"user-tips"}," 新用户可直接登录 ",-1)),ua=U({__name:"index",setup(s){const e=_(!0);let d=E();const c=_(!0),i=_({email:"1978992154@qq.com",ver_code:""}),p=_("获取验证码");let o=120,n=null;F();function w(){e.value=!0}const L=async()=>{try{await d.login(i.value),z.success("登录成功"),await d.queryUserInfo()}catch{z.error("登录失败")}},P=S=>{let r={email:i.value.email,type_code:1};S.preventDefault(),c.value&&(X(r).then(f=>{B.success("验证码发送成功"),c.value=!1,p.value=`${o}s后重发`,n=setInterval(()=>{o--,p.value=`${o}s后重发`,o===0&&(clearInterval(n),o=120,c.value=!0,p.value="获取验证码")},1e3)}).catch(f=>{B.error(f.response.data.msg),c.value=!0}),console.log("get email code"))};return(S,r)=>{const f=v("el-input"),q=v("SvgIcon"),y=v("el-form-item"),H=v("el-button"),M=v("el-form"),R=v("el-dialog");return l(),u("div",da,[a("div",{class:"login_btn"},[a("button",{class:"btn",onClick:w},"登录")]),t(R,{class:"login-dialog",modelValue:e.value,"onUpdate:modelValue":r[2]||(r[2]=x=>e.value=x)},{default:h(()=>[a("div",na,[ia,a("div",la,[oa,a("div",va,[t(M,{model:i.value,class:"form","label-width":"auto"},{default:h(()=>[t(y,{class:"emial"},{default:h(()=>[t(f,{placeholder:"请输入邮箱",modelValue:i.value.email,"onUpdate:modelValue":r[0]||(r[0]=x=>i.value.email=x)},null,8,["modelValue"]),t(q,{name:"email",class:"icon",width:"25",height:"25"})]),_:1}),t(y,{class:"code"},{default:h(()=>[t(f,{placeholder:"请输入验证码",disabled:c.value,modelValue:i.value.ver_code,"onUpdate:modelValue":r[1]||(r[1]=x=>i.value.ver_code=x)},null,8,["disabled","modelValue"]),a("button",{class:"verfi-code",disabled:!c.value,onClick:P},T(p.value),9,ha)]),_:1}),_a,t(y,{class:"btn"},{default:h(()=>[t(H,{type:"primary",onClick:L},{default:h(()=>[O("登陆")]),_:1})]),_:1})]),_:1},8,["model"])]),pa,ra])])]),_:1},8,["modelValue"])])}}}),ga=m(ua,[["__scopeId","data-v-a6cd265a"]]),ma="data:image/svg+xml,%3c?xml%20version='1.0'%20standalone='no'?%3e%3c!DOCTYPE%20svg%20PUBLIC%20'-//W3C//DTD%20SVG%201.1//EN'%20'http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd'%3e%3csvg%20t='1723456942503'%20class='icon'%20viewBox='0%200%201024%201024'%20version='1.1'%20xmlns='http://www.w3.org/2000/svg'%20p-id='992'%20xmlns:xlink='http://www.w3.org/1999/xlink'%20width='200'%20height='200'%3e%3cpath%20d='M468.906667%2057.941333a68.394667%2068.394667%200%200%201%2086.165333%200l399.936%20322.624c30.058667%2024.234667%2035.136%2068.693333%2011.328%2099.328a69.269333%2069.269333%200%200%201-37.802667%2024.768l-3.712%200.810667%200.021334%20361.514667c0%2059.370667-44.48%20108.48-102.528%20113.877333l-4.736%200.362667-4.949334%200.106666H211.370667c-59.84%200-109.056-47.808-112.106667-109.312l-0.106667-5.034666-0.021333-361.514667-2.453333-0.490667c-29.696-6.848-52.16-33.088-53.909334-64.96L42.666667%20436.010667c0-21.610667%209.685333-42.026667%2026.325333-55.466667z%20m46.72%2050.026667a5.76%205.76%200%200%200-7.253334%200L108.864%20431.146667a6.058667%206.058667%200%200%200-2.218667%204.693333c0%203.328%202.624%206.016%205.866667%206.016h18.816c17.536-0.021333%2031.744%2014.506667%2031.744%2032.405333v392.533334l0.064%203.392C164.437333%20896.533333%20185.792%20917.333333%20211.712%20917.333333h599.765333l3.328-0.085333c25.770667-1.322667%2046.101333-23.146667%2046.101334-49.621333V474.282667c0-17.92%2014.208-32.426667%2031.744-32.426667h18.816c1.792%200%203.477333-0.832%204.608-2.261333a6.101333%206.101333%200%200%200-0.96-8.426667z%20m138.133333%20564.693333a32%2032%200%200%201-3.754667%2045.098667C610.858667%20750.890667%20568.106667%20768%20522.666667%20768c-45.44%200-88.192-17.109333-127.36-50.24a32%2032%200%201%201%2041.386666-48.853333C464.704%20692.650667%20493.056%20704%20522.666667%20704c29.589333%200%2057.941333-11.349333%2085.973333-35.093333a32%2032%200%200%201%2045.12%203.754666z'%20fill='%23111111'%20p-id='993'%3e%3c/path%3e%3c/svg%3e",wa="data:image/svg+xml,%3c?xml%20version='1.0'%20standalone='no'?%3e%3c!DOCTYPE%20svg%20PUBLIC%20'-//W3C//DTD%20SVG%201.1//EN'%20'http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd'%3e%3csvg%20t='1723457039167'%20class='icon'%20viewBox='0%200%201024%201024'%20version='1.1'%20xmlns='http://www.w3.org/2000/svg'%20p-id='4295'%20xmlns:xlink='http://www.w3.org/1999/xlink'%20width='200'%20height='200'%3e%3cpath%20d='M796.202667%20101.973333a98.218667%2098.218667%200%200%201%20125.866666%20125.738667l-211.285333%20632.746667v0.085333A98.048%2098.048%200%200%201%20640%20925.269333a96.981333%2096.981333%200%200%201-92.16-26.325333l-112.085333-111.488-117.504%2060.714667a32%2032%200%200%201-46.677334-29.269334l5.333334-204.8a32%2032%200%200%201%2011.52-23.765333l503.765333-419.584-608.853333%20202.965333h-0.042667a34.133333%2034.133333%200%200%200-22.485333%2026.197334l-0.128%200.725333a32.853333%2032.853333%200%200%200%208.96%2029.952l70.101333%2069.973333a32%2032%200%201%201-45.226667%2045.312l-70.058666-69.973333a96.981333%2096.981333%200%200%201-26.538667-87.765333A98.090667%2098.090667%200%200%201%20162.688%20313.173333l0.213333-0.085333%20633.301334-211.157333z%20m67.797333%2092.245334L340.48%20630.272l-3.498667%20136.192%2090.112-46.549333a32%2032%200%200%201%2037.248%205.717333l129.109334%20128.426667a32.981333%2032.981333%200%200%200%2031.488%208.96l0.384-0.085334a34.133333%2034.133333%200%200%200%2024.704-22.528v-0.085333l211.498666-633.301333a34.048%2034.048%200%200%200%202.474667-12.8z'%20fill='%231F2322'%20p-id='4296'%3e%3c/path%3e%3c/svg%3e",fa="data:image/svg+xml,%3c?xml%20version='1.0'%20standalone='no'?%3e%3c!DOCTYPE%20svg%20PUBLIC%20'-//W3C//DTD%20SVG%201.1//EN'%20'http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd'%3e%3csvg%20t='1723449260981'%20class='icon'%20viewBox='0%200%201024%201024'%20version='1.1'%20xmlns='http://www.w3.org/2000/svg'%20p-id='1143'%20xmlns:xlink='http://www.w3.org/1999/xlink'%20width='200'%20height='200'%3e%3cpath%20d='M544%20161.536a330.666667%20330.666667%200%200%201%20298.666667%20329.130667h-0.341334c0.213333%201.493333%200.341333%202.986667%200.341334%204.565333v219.434667h39.68a32%2032%200%200%201%200%2064h-212.053334a160%20160%200%200%201-316.586666%200H141.909333a32%2032%200%201%201%200-64h39.424v-219.434667c0-1.578667%200.128-3.072%200.341334-4.565333H181.333333a330.666667%20330.666667%200%200%201%20298.666667-329.130667V128a32%2032%200%201%201%2064%200v33.536z%20m-298.666667%20553.130667h533.333334v-219.434667c0-1.578667%200.128-3.072%200.341333-4.565333h-0.341333a266.666667%20266.666667%200%201%200-533.333334%200h-0.341333c0.213333%201.493333%200.341333%202.986667%200.341333%204.565333v219.434667z%20m359.765334%2064H418.901333a96%2096%200%200%200%20186.197334%200z'%20fill='%23111111'%20p-id='1144'%3e%3c/path%3e%3c/svg%3e",xa=s=>(V("data-v-f0f1f535"),s=s(),I(),s),$a={class:"channel_container"},Ca={class:"icon"},ba=["src"],ka={key:0,class:"count"},Va={class:"title"},Ia={class:"icon"},Sa=["src"],ya=xa(()=>a("div",{class:"title"},"我",-1)),Da={__name:"index",setup(s){const e=E(),d=_(e.getUserInfo),c=_([{name:"home",label:"发现",icon:ma,to:"/home",count:0},{name:"publish",label:"发布",icon:wa,to:"/publish",count:0},{name:"notices",label:"通知",icon:fa,to:"/notices",count:4}]),i=_("home");return(p,o)=>(l(),u("div",$a,[(l(!0),u(W,null,Y(c.value,(n,w)=>(l(),$(g(G),{class:j(["item",{active:i.value===n.name}]),key:w,onClick:L=>i.value=n.name,to:n.to},{default:h(()=>[a("div",Ca,[a("img",{src:n.icon,alt:""},null,8,ba),n.count>0?(l(),u("div",ka,T(n.count),1)):k("",!0)]),a("div",Va,T(n.label),1)]),_:2},1032,["class","onClick","to"]))),128)),g(e).isLogin()?(l(),$(g(G),{key:0,class:"item",to:`/user/profile/${d.value.globalNumber}`},{default:h(()=>[a("div",Ia,[a("img",{src:d.value.avatar?d.value.avatar:"../../assets/imgs/avatar.jpeg",class:"avatar",alt:""},null,8,Sa)]),ya]),_:1},8,["to"])):k("",!0),g(e).isLogin()?k("",!0):(l(),$(ga,{key:1}))]))}},Ta=m(Da,[["__scopeId","data-v-f0f1f535"]]),b=s=>(V("data-v-789d34db"),s=s(),I(),s),Ea={class:"card"},La=b(()=>a("div",{class:"title"},"马上登录即可",-1)),za={class:"line_container"},Ba=b(()=>a("span",null,"刷到更懂你的优质内容",-1)),Ga={class:"line_container"},Na=b(()=>a("span",null,"搜索最新种草、拔草信息",-1)),Ua={class:"line_container"},Pa=b(()=>a("span",null,"查看收藏、点赞的笔记",-1)),qa={class:"line_container"},Ha=b(()=>a("span",null,"与他人更好地互动、交流",-1)),Ma={__name:"index",setup(s){return(e,d)=>{const c=v("svg-icon");return l(),u("div",Ea,[La,a("div",za,[t(c,{name:"nice",class:"icon",width:"16",height:"16"}),Ba]),a("div",Ga,[t(c,{name:"zhongcao",class:"icon",width:"16",height:"16"}),Na]),a("div",Ua,[t(c,{name:"star",class:"icon",width:"16",height:"16"}),Pa]),a("div",qa,[t(c,{name:"talk",class:"icon",width:"16",height:"16"}),Ha])])}}},Ra=m(Ma,[["__scopeId","data-v-789d34db"]]),Fa={class:"search_box"},Oa={class:"input_button"},Wa={class:"sug-container-wrapper sug_pad"},Ya=K('<div class="sug-container" data-v-65ca3d62><div class="sug-box" data-v-65ca3d62><div class="sug-wrapper" data-v-65ca3d62><div class="sug-item active" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>折</span><span class="" data-v-65ca3d62>纸</span><span class="" data-v-65ca3d62>教</span><span class="" data-v-65ca3d62>程</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>球</span><span class="" data-v-65ca3d62>兰</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>蛋</span><span class="" data-v-65ca3d62>糕</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>湖</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>函</span><span class="" data-v-65ca3d62>数</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>岛</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>脸</span><span class="" data-v-65ca3d62>发</span><span class="" data-v-65ca3d62>型</span><span class="" data-v-65ca3d62>男</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>树</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>水</span><span class="" data-v-65ca3d62>果</span><span class="" data-v-65ca3d62>礼</span><span class="" data-v-65ca3d62>盒</span></div></div></div><div class="sug-item" data-v-65ca3d62><div data-v-65ca3d62><div data-v-65ca3d62><span class="highlight" data-v-65ca3d62>心</span><span class="highlight" data-v-65ca3d62>形</span><span class="" data-v-65ca3d62>慕</span><span class="" data-v-65ca3d62>斯</span></div></div></div></div></div><div data-v-65ca3d62></div></div>',1),ja=[Ya],Aa={__name:"index",setup(s){const e=_(""),d=_(!1);A(e,p=>{d.value=p.length>0});function c(){e.value=""}function i(){console.log("search")}return(p,o)=>{const n=v("svg-icon");return l(),u("div",Fa,[D(a("input",{type:"text","onUpdate:modelValue":o[0]||(o[0]=w=>e.value=w),class:"search_input",placeholder:"搜索小黄书"},null,512),[[J,e.value]]),a("div",Oa,[D(a("div",{class:"close_icon",onClick:c},[t(n,{name:"close",width:"16",height:"16",color:"#000"})],512),[[N,d.value]]),a("div",{class:"search_icon",onClick:i},[t(n,{name:"search",width:"16",height:"16",color:"#000"})])]),D(a("div",Wa,ja,512),[[N,d.value]])])}}},Ja=m(Aa,[["__scopeId","data-v-65ca3d62"]]),Ka={class:"layout_container"},Qa={class:"layout_slide"},Xa={class:"sider_main"},Za={class:"layout_header"},a3={class:"layout_main"},s3=U({__name:"index",setup(s){const e=E();return(d,c)=>(l(),u("div",Ka,[a("div",Qa,[t(ca),a("div",Xa,[t(Ta),g(e).getToken?k("",!0):(l(),$(Ra,{key:0}))])]),a("div",Za,[t(Ja)]),a("div",a3,[t(g(Q))])]))}}),d3=m(s3,[["__scopeId","data-v-ce204c9b"]]);export{d3 as default};
