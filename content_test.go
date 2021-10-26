package fsfire

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkContentPrefixClear(t *testing.T) {
	grids := []struct {
		data               []byte
		mark               string
		containsMarkString bool
		expected           string
	}{
		{
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/textfiles/huaji.txt")
				assert.NoError(t, err)

				return data
			}(),
			mark:               "我哀！为这些在疫情中逝去的亡灵！",
			containsMarkString: true,
			expected: `我哀！为这些在疫情中逝去的亡灵！
我哭！为这人间遭受的灾难！
我痛！在疫情中以身殉职的李文亮烈士！平凡而普通的你，凭着职业的敏感与做人的良知，第一时间发布了病毒预警；正直而坚强的你，顶着巨大的压力，继续在一线工作，恪尽职守；肩负责任和良知的你，当不幸感染了病毒，已经住进了重症监护室，还对记者表示，康复之后仍然要奔赴一线。英雄的生命最终被该诅咒的病毒夺走，但却牵动了亿万人的心。你的微博，已经成为了中国互联网的哭墙，每天有成千上万的人们向你问候，向你致哀。李医生，我们的李医生，你在天堂看到了吗？听到了吗？
漫天飞舞的樱花啊，飘飘洒洒，纷纷扬扬，每一瓣分明就是破碎了的心瓣，滴着血，滴着血，带着生者无限的哀思，飘过了清明时节的天空……
在水一方写于2020年清明节前夕
`,
		},
		{
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/textfiles/huaji.txt")
				assert.NoError(t, err)

				return data
			}(),
			mark:               "我哀！为这些在疫情中逝去的亡灵！",
			containsMarkString: false,
			expected: `
我哭！为这人间遭受的灾难！
我痛！在疫情中以身殉职的李文亮烈士！平凡而普通的你，凭着职业的敏感与做人的良知，第一时间发布了病毒预警；正直而坚强的你，顶着巨大的压力，继续在一线工作，恪尽职守；肩负责任和良知的你，当不幸感染了病毒，已经住进了重症监护室，还对记者表示，康复之后仍然要奔赴一线。英雄的生命最终被该诅咒的病毒夺走，但却牵动了亿万人的心。你的微博，已经成为了中国互联网的哭墙，每天有成千上万的人们向你问候，向你致哀。李医生，我们的李医生，你在天堂看到了吗？听到了吗？
漫天飞舞的樱花啊，飘飘洒洒，纷纷扬扬，每一瓣分明就是破碎了的心瓣，滴着血，滴着血，带着生者无限的哀思，飘过了清明时节的天空……
在水一方写于2020年清明节前夕
`,
		},
		{
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/textfiles/huaji.txt")
				assert.NoError(t, err)

				return data
			}(),
			expected: `这一年的春天来得凝重，来得孤寂，以至于武汉的樱花铺天盖地像往年一样深情地向大地报告春讯时，我受创的泪腺，依然滴下冰冷的泪水；这一年的清明节来得如此匆匆，来得如此悄然，以至于因哀痛堵满的心，平添了窒息的痛苦。
一场突然而来的新冠肺炎病毒，于去年年底在武汉疯狂肆虐，无情地让数以万计的人们承受着感染病毒的残酷折磨，数以千计的生命转瞬即逝，许许多多的家庭支离破碎。武汉因此封城，全国为此按下了暂停键。我们，正赴一场罕见的国难。
多么希望平生见证的这段悲痛历史可以改写；多么希望那些逝去的鲜活灵魂可以重生；多么希望，在这伤春之际，在这清明时节，以一个幸存者的身份，为那些无辜踏上黄泉路的遇难者我的同胞凭吊。
于是，我拿起了画笔，画下了几朵心中的花。如果非要问，花开为谁祭 ，花谢为谁泣 ？我想，我要为人生阶段的幼年、青年、中年、老年逝者致哀。
　　掬一瓣花香，祭幼年的你！孩子，你来到这个世界，犹如是枝头刚伸展的嫰芽，还沒来得及绽放，就被一阵狂风给刮走了。我多么希望你幼小的心灵，能多一天感受父母的宠爱；我多么希望你纯真无邪的双眼，能再次饱览这春暖花开的美景。我多么希望你一天一天健康成长，年复一年地度过平凡人的一生：读书、就业、恋爱、成家。就这样目睹你生命的延续。人间所有的喜与乐，苦与痛，你都可以在漫长的一生里，慢慢享受，慢慢感悟，慢慢去领略生命的美好。孩子，亲爱的孩子，你的突然夭折，我宁愿相信这是天使太怜惜你了，太心疼你了，她不忍心病毒在你幼小的身子上继续狂魔乱舞。孩子，你可以放心地跟随天使姐姐走吧，她会牵引你走向天堂。因为，天堂里没有病毒……
掬一瓣花香，祭青年的你！青春，是生命的高光；青春，是人生之花刚刚绽放。而青年的你，还沒来得及唱完这首青春之歌，还没来得及让生命之花怒放，蓬勃的生命竟然会因病毒的一时侵袭戛然而止。瘁不及防的打击，摧毁了父母这辈子最大的希望，也扼杀了那一场场来不及恋的爱。当我看到那位还沒毕业的女博士生，在遗书里痛悔没有把初吻献给心爱的人时，我感到的是撕心裂肺的痛！可恶的病毒，你不仅夺走了如朝霞般美丽的生命，你还夺走了一个人在最美好的年华里应该享有的爱与被爱的权利。人间是多么的无情和残酷，我唯有祈望在天堂里，青年的你可以续写青春之歌……
掬一瓣花香，祭中年的你！我知道在所有的逝者里，中年的你走得异常艰难。人生历程里的这个阶段，上有老下有小，你是家的天，你是亲人的依靠。你深知肩上负载的责任，你感知自己在亲人心中的重量。平时，再苦再累你也假装若无其事。如今，面对死神的召唤，你显得多么地悲怆和哀痛。求生的欲望经久地在脑海里盘旋，你渴望奇迹降临，生机重现。当你痛苦闭上双眼时，眼角留下了两行绝望的泪水。唯愿在天堂里，你可以卸下所有的包袱，轻装前行……
掬一瓣花香，祭老年的你！在本该颐养天年以及含饴弄孙的晚年，突遭杀身之祸。从天而降的这场新冠病毒，就是一场噩梦的开始。它让患者临死前经历一场人间酷刑的折磨。你会吃不下，干咳，浑身酸痛无力，呼吸困难。更可怕的是，这种病是死于呼吸衰竭，并没造成身体上的神志障碍和意识丧失。所以，你是看着自己痛苦挣扎，看着自己活活窒息。天哪！这种地狱般的惨烈折磨，让耄耋之年的你临终发生，在你羸弱的身躯恣暴，该是多么地悲惨与恐怖！当你咽下最后一口气的时候，我相信你已从痛苦的挣扎中解脱，天堂的大门已为你敞开……
我哀！为这些在疫情中逝去的亡灵！
我哭！为这人间遭受的灾难！
我痛！在疫情中以身殉职的李文亮烈士！平凡而普通的你，凭着职业的敏感与做人的良知，第一时间发布了病毒预警；正直而坚强的你，顶着巨大的压力，继续在一线工作，恪尽职守；肩负责任和良知的你，当不幸感染了病毒，已经住进了重症监护室，还对记者表示，康复之后仍然要奔赴一线。英雄的生命最终被该诅咒的病毒夺走，但却牵动了亿万人的心。你的微博，已经成为了中国互联网的哭墙，每天有成千上万的人们向你问候，向你致哀。李医生，我们的李医生，你在天堂看到了吗？听到了吗？
漫天飞舞的樱花啊，飘飘洒洒，纷纷扬扬，每一瓣分明就是破碎了的心瓣，滴着血，滴着血，带着生者无限的哀思，飘过了清明时节的天空……
在水一方写于2020年清明节前夕
`,
		},
	}

	for _, grid := range grids {
		actual := MarkContentPrefixClear(grid.data, grid.mark, WithSpecificContainsMarkString(grid.containsMarkString))
		assert.Equal(t, grid.expected, string(actual))
	}
}

func TestMarkContentSuffixClear(t *testing.T) {
	grids := []struct {
		data               []byte
		mark               string
		containsMarkString bool
		expected           string
	}{
		{
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/textfiles/huaji.txt")
				assert.NoError(t, err)

				return data
			}(),
			mark:               "我哀！为这些在疫情中逝去的亡灵！",
			containsMarkString: false,
			expected: `这一年的春天来得凝重，来得孤寂，以至于武汉的樱花铺天盖地像往年一样深情地向大地报告春讯时，我受创的泪腺，依然滴下冰冷的泪水；这一年的清明节来得如此匆匆，来得如此悄然，以至于因哀痛堵满的心，平添了窒息的痛苦。
一场突然而来的新冠肺炎病毒，于去年年底在武汉疯狂肆虐，无情地让数以万计的人们承受着感染病毒的残酷折磨，数以千计的生命转瞬即逝，许许多多的家庭支离破碎。武汉因此封城，全国为此按下了暂停键。我们，正赴一场罕见的国难。
多么希望平生见证的这段悲痛历史可以改写；多么希望那些逝去的鲜活灵魂可以重生；多么希望，在这伤春之际，在这清明时节，以一个幸存者的身份，为那些无辜踏上黄泉路的遇难者我的同胞凭吊。
于是，我拿起了画笔，画下了几朵心中的花。如果非要问，花开为谁祭 ，花谢为谁泣 ？我想，我要为人生阶段的幼年、青年、中年、老年逝者致哀。
　　掬一瓣花香，祭幼年的你！孩子，你来到这个世界，犹如是枝头刚伸展的嫰芽，还沒来得及绽放，就被一阵狂风给刮走了。我多么希望你幼小的心灵，能多一天感受父母的宠爱；我多么希望你纯真无邪的双眼，能再次饱览这春暖花开的美景。我多么希望你一天一天健康成长，年复一年地度过平凡人的一生：读书、就业、恋爱、成家。就这样目睹你生命的延续。人间所有的喜与乐，苦与痛，你都可以在漫长的一生里，慢慢享受，慢慢感悟，慢慢去领略生命的美好。孩子，亲爱的孩子，你的突然夭折，我宁愿相信这是天使太怜惜你了，太心疼你了，她不忍心病毒在你幼小的身子上继续狂魔乱舞。孩子，你可以放心地跟随天使姐姐走吧，她会牵引你走向天堂。因为，天堂里没有病毒……
掬一瓣花香，祭青年的你！青春，是生命的高光；青春，是人生之花刚刚绽放。而青年的你，还沒来得及唱完这首青春之歌，还没来得及让生命之花怒放，蓬勃的生命竟然会因病毒的一时侵袭戛然而止。瘁不及防的打击，摧毁了父母这辈子最大的希望，也扼杀了那一场场来不及恋的爱。当我看到那位还沒毕业的女博士生，在遗书里痛悔没有把初吻献给心爱的人时，我感到的是撕心裂肺的痛！可恶的病毒，你不仅夺走了如朝霞般美丽的生命，你还夺走了一个人在最美好的年华里应该享有的爱与被爱的权利。人间是多么的无情和残酷，我唯有祈望在天堂里，青年的你可以续写青春之歌……
掬一瓣花香，祭中年的你！我知道在所有的逝者里，中年的你走得异常艰难。人生历程里的这个阶段，上有老下有小，你是家的天，你是亲人的依靠。你深知肩上负载的责任，你感知自己在亲人心中的重量。平时，再苦再累你也假装若无其事。如今，面对死神的召唤，你显得多么地悲怆和哀痛。求生的欲望经久地在脑海里盘旋，你渴望奇迹降临，生机重现。当你痛苦闭上双眼时，眼角留下了两行绝望的泪水。唯愿在天堂里，你可以卸下所有的包袱，轻装前行……
掬一瓣花香，祭老年的你！在本该颐养天年以及含饴弄孙的晚年，突遭杀身之祸。从天而降的这场新冠病毒，就是一场噩梦的开始。它让患者临死前经历一场人间酷刑的折磨。你会吃不下，干咳，浑身酸痛无力，呼吸困难。更可怕的是，这种病是死于呼吸衰竭，并没造成身体上的神志障碍和意识丧失。所以，你是看着自己痛苦挣扎，看着自己活活窒息。天哪！这种地狱般的惨烈折磨，让耄耋之年的你临终发生，在你羸弱的身躯恣暴，该是多么地悲惨与恐怖！当你咽下最后一口气的时候，我相信你已从痛苦的挣扎中解脱，天堂的大门已为你敞开……
`,
		},

		{
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/textfiles/huaji.txt")
				assert.NoError(t, err)

				return data
			}(),
			mark:               "我哀！为这些在疫情中逝去的亡灵！",
			containsMarkString: true,
			expected: `这一年的春天来得凝重，来得孤寂，以至于武汉的樱花铺天盖地像往年一样深情地向大地报告春讯时，我受创的泪腺，依然滴下冰冷的泪水；这一年的清明节来得如此匆匆，来得如此悄然，以至于因哀痛堵满的心，平添了窒息的痛苦。
一场突然而来的新冠肺炎病毒，于去年年底在武汉疯狂肆虐，无情地让数以万计的人们承受着感染病毒的残酷折磨，数以千计的生命转瞬即逝，许许多多的家庭支离破碎。武汉因此封城，全国为此按下了暂停键。我们，正赴一场罕见的国难。
多么希望平生见证的这段悲痛历史可以改写；多么希望那些逝去的鲜活灵魂可以重生；多么希望，在这伤春之际，在这清明时节，以一个幸存者的身份，为那些无辜踏上黄泉路的遇难者我的同胞凭吊。
于是，我拿起了画笔，画下了几朵心中的花。如果非要问，花开为谁祭 ，花谢为谁泣 ？我想，我要为人生阶段的幼年、青年、中年、老年逝者致哀。
　　掬一瓣花香，祭幼年的你！孩子，你来到这个世界，犹如是枝头刚伸展的嫰芽，还沒来得及绽放，就被一阵狂风给刮走了。我多么希望你幼小的心灵，能多一天感受父母的宠爱；我多么希望你纯真无邪的双眼，能再次饱览这春暖花开的美景。我多么希望你一天一天健康成长，年复一年地度过平凡人的一生：读书、就业、恋爱、成家。就这样目睹你生命的延续。人间所有的喜与乐，苦与痛，你都可以在漫长的一生里，慢慢享受，慢慢感悟，慢慢去领略生命的美好。孩子，亲爱的孩子，你的突然夭折，我宁愿相信这是天使太怜惜你了，太心疼你了，她不忍心病毒在你幼小的身子上继续狂魔乱舞。孩子，你可以放心地跟随天使姐姐走吧，她会牵引你走向天堂。因为，天堂里没有病毒……
掬一瓣花香，祭青年的你！青春，是生命的高光；青春，是人生之花刚刚绽放。而青年的你，还沒来得及唱完这首青春之歌，还没来得及让生命之花怒放，蓬勃的生命竟然会因病毒的一时侵袭戛然而止。瘁不及防的打击，摧毁了父母这辈子最大的希望，也扼杀了那一场场来不及恋的爱。当我看到那位还沒毕业的女博士生，在遗书里痛悔没有把初吻献给心爱的人时，我感到的是撕心裂肺的痛！可恶的病毒，你不仅夺走了如朝霞般美丽的生命，你还夺走了一个人在最美好的年华里应该享有的爱与被爱的权利。人间是多么的无情和残酷，我唯有祈望在天堂里，青年的你可以续写青春之歌……
掬一瓣花香，祭中年的你！我知道在所有的逝者里，中年的你走得异常艰难。人生历程里的这个阶段，上有老下有小，你是家的天，你是亲人的依靠。你深知肩上负载的责任，你感知自己在亲人心中的重量。平时，再苦再累你也假装若无其事。如今，面对死神的召唤，你显得多么地悲怆和哀痛。求生的欲望经久地在脑海里盘旋，你渴望奇迹降临，生机重现。当你痛苦闭上双眼时，眼角留下了两行绝望的泪水。唯愿在天堂里，你可以卸下所有的包袱，轻装前行……
掬一瓣花香，祭老年的你！在本该颐养天年以及含饴弄孙的晚年，突遭杀身之祸。从天而降的这场新冠病毒，就是一场噩梦的开始。它让患者临死前经历一场人间酷刑的折磨。你会吃不下，干咳，浑身酸痛无力，呼吸困难。更可怕的是，这种病是死于呼吸衰竭，并没造成身体上的神志障碍和意识丧失。所以，你是看着自己痛苦挣扎，看着自己活活窒息。天哪！这种地狱般的惨烈折磨，让耄耋之年的你临终发生，在你羸弱的身躯恣暴，该是多么地悲惨与恐怖！当你咽下最后一口气的时候，我相信你已从痛苦的挣扎中解脱，天堂的大门已为你敞开……
我哀！为这些在疫情中逝去的亡灵！`,
		},

		{
			data: func() []byte {
				data, err := ioutil.ReadFile("test/data/assets/textfiles/huaji.txt")
				assert.NoError(t, err)

				return data
			}(),
			expected: `这一年的春天来得凝重，来得孤寂，以至于武汉的樱花铺天盖地像往年一样深情地向大地报告春讯时，我受创的泪腺，依然滴下冰冷的泪水；这一年的清明节来得如此匆匆，来得如此悄然，以至于因哀痛堵满的心，平添了窒息的痛苦。
一场突然而来的新冠肺炎病毒，于去年年底在武汉疯狂肆虐，无情地让数以万计的人们承受着感染病毒的残酷折磨，数以千计的生命转瞬即逝，许许多多的家庭支离破碎。武汉因此封城，全国为此按下了暂停键。我们，正赴一场罕见的国难。
多么希望平生见证的这段悲痛历史可以改写；多么希望那些逝去的鲜活灵魂可以重生；多么希望，在这伤春之际，在这清明时节，以一个幸存者的身份，为那些无辜踏上黄泉路的遇难者我的同胞凭吊。
于是，我拿起了画笔，画下了几朵心中的花。如果非要问，花开为谁祭 ，花谢为谁泣 ？我想，我要为人生阶段的幼年、青年、中年、老年逝者致哀。
　　掬一瓣花香，祭幼年的你！孩子，你来到这个世界，犹如是枝头刚伸展的嫰芽，还沒来得及绽放，就被一阵狂风给刮走了。我多么希望你幼小的心灵，能多一天感受父母的宠爱；我多么希望你纯真无邪的双眼，能再次饱览这春暖花开的美景。我多么希望你一天一天健康成长，年复一年地度过平凡人的一生：读书、就业、恋爱、成家。就这样目睹你生命的延续。人间所有的喜与乐，苦与痛，你都可以在漫长的一生里，慢慢享受，慢慢感悟，慢慢去领略生命的美好。孩子，亲爱的孩子，你的突然夭折，我宁愿相信这是天使太怜惜你了，太心疼你了，她不忍心病毒在你幼小的身子上继续狂魔乱舞。孩子，你可以放心地跟随天使姐姐走吧，她会牵引你走向天堂。因为，天堂里没有病毒……
掬一瓣花香，祭青年的你！青春，是生命的高光；青春，是人生之花刚刚绽放。而青年的你，还沒来得及唱完这首青春之歌，还没来得及让生命之花怒放，蓬勃的生命竟然会因病毒的一时侵袭戛然而止。瘁不及防的打击，摧毁了父母这辈子最大的希望，也扼杀了那一场场来不及恋的爱。当我看到那位还沒毕业的女博士生，在遗书里痛悔没有把初吻献给心爱的人时，我感到的是撕心裂肺的痛！可恶的病毒，你不仅夺走了如朝霞般美丽的生命，你还夺走了一个人在最美好的年华里应该享有的爱与被爱的权利。人间是多么的无情和残酷，我唯有祈望在天堂里，青年的你可以续写青春之歌……
掬一瓣花香，祭中年的你！我知道在所有的逝者里，中年的你走得异常艰难。人生历程里的这个阶段，上有老下有小，你是家的天，你是亲人的依靠。你深知肩上负载的责任，你感知自己在亲人心中的重量。平时，再苦再累你也假装若无其事。如今，面对死神的召唤，你显得多么地悲怆和哀痛。求生的欲望经久地在脑海里盘旋，你渴望奇迹降临，生机重现。当你痛苦闭上双眼时，眼角留下了两行绝望的泪水。唯愿在天堂里，你可以卸下所有的包袱，轻装前行……
掬一瓣花香，祭老年的你！在本该颐养天年以及含饴弄孙的晚年，突遭杀身之祸。从天而降的这场新冠病毒，就是一场噩梦的开始。它让患者临死前经历一场人间酷刑的折磨。你会吃不下，干咳，浑身酸痛无力，呼吸困难。更可怕的是，这种病是死于呼吸衰竭，并没造成身体上的神志障碍和意识丧失。所以，你是看着自己痛苦挣扎，看着自己活活窒息。天哪！这种地狱般的惨烈折磨，让耄耋之年的你临终发生，在你羸弱的身躯恣暴，该是多么地悲惨与恐怖！当你咽下最后一口气的时候，我相信你已从痛苦的挣扎中解脱，天堂的大门已为你敞开……
我哀！为这些在疫情中逝去的亡灵！
我哭！为这人间遭受的灾难！
我痛！在疫情中以身殉职的李文亮烈士！平凡而普通的你，凭着职业的敏感与做人的良知，第一时间发布了病毒预警；正直而坚强的你，顶着巨大的压力，继续在一线工作，恪尽职守；肩负责任和良知的你，当不幸感染了病毒，已经住进了重症监护室，还对记者表示，康复之后仍然要奔赴一线。英雄的生命最终被该诅咒的病毒夺走，但却牵动了亿万人的心。你的微博，已经成为了中国互联网的哭墙，每天有成千上万的人们向你问候，向你致哀。李医生，我们的李医生，你在天堂看到了吗？听到了吗？
漫天飞舞的樱花啊，飘飘洒洒，纷纷扬扬，每一瓣分明就是破碎了的心瓣，滴着血，滴着血，带着生者无限的哀思，飘过了清明时节的天空……
在水一方写于2020年清明节前夕
`,
		},
	}

	for _, grid := range grids {
		actual := MarkContentSuffixClear(grid.data, grid.mark, WithSpecificContainsMarkString(grid.containsMarkString))
		assert.Equal(t, grid.expected, string(actual))
	}
}