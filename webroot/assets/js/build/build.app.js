({
	optimize: "none",
	//optimize: "uglify2",
	paths: {
		'util' : 'mylib/util',
		'ajax' : 'mylib/ajax',
		'tool' : 'mylib/tool',
		'accto' : 'mylib/accto',
		'preparecookie' : 'mylib/preparecookie',
		'searchbar' : 'mylib/vue-searchbar',

		'notepad' : 'mylib/vue/notepad',
		'areas' : 'mylib/vue/comp/vue-areas',
		'contents' : 'mylib/vue/comp/vue-contents',
		'subforms' : 'mylib/vue/comp/vue-subforms',
		'tags' : 'mylib/vue/comp/vue-tags',

    'compressor' : 'empty:',
    'echarts' : 'empty:',
    'vue' : 'empty:'
	},

	baseUrl : "../dev",
	removeCombined : true,
	modules: [
		{
			name : 'page/home/Default',
			create: false
		},
		{
			name : 'page/home/Login',
			create: false
		},
		{
			name : 'page/home/Cell',
			create: false
		},
		{
			name : 'page/home/Cells',
			create: false
		}
	],

	dir : '../pro'
})
