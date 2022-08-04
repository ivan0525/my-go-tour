import{_ as s,c as n,o as a,a as e}from"./app.8ab4ad29.js";const g=JSON.parse('{"title":"\u4F9D\u8D56\u7BA1\u7406","description":"","frontmatter":{},"headers":[{"level":2,"title":"1. go mod\u547D\u4EE4","slug":"_1-go-mod\u547D\u4EE4"}],"relativePath":"basic/dependency.md"}'),l={name:"basic/dependency.md"},o=e(`<h1 id="\u4F9D\u8D56\u7BA1\u7406" tabindex="-1">\u4F9D\u8D56\u7BA1\u7406 <a class="header-anchor" href="#\u4F9D\u8D56\u7BA1\u7406" aria-hidden="true">#</a></h1><h2 id="_1-go-mod\u547D\u4EE4" tabindex="-1">1. <code>go mod</code>\u547D\u4EE4 <a class="header-anchor" href="#_1-go-mod\u547D\u4EE4" aria-hidden="true">#</a></h2><p>\u5E38\u7528\u7684<code>go mod</code>\u547D\u4EE4</p><div class="language-bash"><span class="copy"></span><pre><code><span class="line"><span style="color:#676E95;font-style:italic;"># \u4E0B\u8F7D\u4F9D\u8D56\u7684module\u5230\u672C\u5730cache\uFF08\u9ED8\u8BA4\u4E3A$GOPATH/pkg/mod\u76EE\u5F55\uFF09</span></span>
<span class="line"><span style="color:#A6ACCD;">go mod download </span></span>
<span class="line"></span>
<span class="line"><span style="color:#676E95;font-style:italic;"># \u7F16\u8F91go.mod\u6587\u4EF6</span></span>
<span class="line"><span style="color:#A6ACCD;">go mod edit</span></span>
<span class="line"></span>
<span class="line"><span style="color:#676E95;font-style:italic;"># \u6253\u5370\u6A21\u5757\u4F9D\u8D56\u56FE</span></span>
<span class="line"><span style="color:#A6ACCD;">go mod graph</span></span>
<span class="line"></span>
<span class="line"><span style="color:#676E95;font-style:italic;"># \u521D\u59CB\u5316\u5F53\u524D\u6587\u4EF6\u5939\uFF0C\u521B\u5EFAgo.mod</span></span>
<span class="line"><span style="color:#A6ACCD;">go mod init</span></span>
<span class="line"></span>
<span class="line"><span style="color:#676E95;font-style:italic;"># \u589E\u52A0\u7F3A\u5C11\u7684module\uFF0C\u5220\u9664\u65E0\u7528\u7684module</span></span>
<span class="line"><span style="color:#A6ACCD;">go mod tidy</span></span>
<span class="line"></span>
<span class="line"><span style="color:#676E95;font-style:italic;"># \u5C06\u4F9D\u8D56\u590D\u5236\u5230vendor\u4E0B</span></span>
<span class="line"><span style="color:#A6ACCD;">go mod vendor</span></span>
<span class="line"></span>
<span class="line"><span style="color:#676E95;font-style:italic;"># \u6821\u9A8C\u4F9D\u8D56</span></span>
<span class="line"><span style="color:#A6ACCD;">go mod verify</span></span>
<span class="line"></span>
<span class="line"><span style="color:#676E95;font-style:italic;"># \u89E3\u91CA\u4E3A\u4EC0\u4E48\u9700\u8981\u4F9D\u8D56</span></span>
<span class="line"><span style="color:#A6ACCD;">go mod why</span></span>
<span class="line"></span></code></pre></div>`,4),p=[o];function c(t,i,d,r,y,_){return a(),n("div",null,p)}var h=s(l,[["render",c]]);export{g as __pageData,h as default};
