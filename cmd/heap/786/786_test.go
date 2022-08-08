package _786

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	assert.True(t, reflect.DeepEqual([]int{2, 5}, kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3)))
	assert.True(t, reflect.DeepEqual([]int{1, 7}, kthSmallestPrimeFraction([]int{1, 7}, 1)))
}

func Test2(t *testing.T) {
	assert.True(t, reflect.DeepEqual([]int{29879, 29881}, kthSmallestPrimeFraction([]int{1, 3, 7, 23, 59, 97, 101, 103, 127, 137, 149, 173, 191, 197, 199, 227, 233, 281, 283, 307, 313, 347, 353, 367, 373, 401, 419, 421, 449, 457, 467, 521, 541, 563, 577, 593, 599, 607, 617, 631, 647, 659, 661, 727, 751, 769, 797, 821, 827, 881, 911, 941, 991, 1009, 1039, 1063, 1069, 1093, 1103, 1109, 1129, 1171, 1213, 1217, 1277, 1291, 1307, 1319, 1321, 1367, 1433, 1439, 1451, 1471, 1481, 1571, 1579, 1597, 1607, 1619, 1627, 1699, 1747, 1753, 1759, 1867, 1901, 1933, 1951, 1987, 1993, 2053, 2069, 2083, 2089, 2113, 2131, 2143, 2153, 2239, 2309, 2311, 2347, 2377, 2417, 2441, 2551, 2557, 2579, 2657, 2677, 2683, 2687, 2693, 2707, 2719, 2753, 2789, 2797, 2803, 2857, 2887, 2903, 2917, 2939, 2963, 2999, 3061, 3083, 3089, 3163, 3203, 3217, 3253, 3271, 3299, 3323, 3457, 3461, 3469, 3499, 3559, 3623, 3637, 3671, 3709, 3727, 3733, 3739, 3767, 3779, 3821, 3851, 3853, 3863, 3889, 3917, 3929, 3931, 3967, 4019, 4079, 4099, 4127, 4159, 4229, 4241, 4261, 4273, 4327, 4339, 4397, 4447, 4451, 4493, 4513, 4519, 4549, 4679, 4703, 4783, 4787, 4799, 4861, 4909, 4919, 4933, 4943, 4973, 5011, 5077, 5081, 5189, 5309, 5347, 5387, 5413, 5417, 5449, 5483, 5501, 5507, 5527, 5557, 5563, 5641, 5653, 5711, 5741, 5743, 5749, 5783, 5791, 5939, 6043, 6053, 6067, 6073, 6091, 6113, 6203, 6211, 6217, 6271, 6277, 6301, 6311, 6323, 6353, 6389, 6473, 6673, 6689, 6761, 6779, 6857, 6863, 6907, 6947, 6949, 6967, 6977, 6991, 7019, 7043, 7057, 7129, 7159, 7193, 7213, 7229, 7243, 7253, 7307, 7309, 7331, 7333, 7369, 7451, 7459, 7481, 7487, 7517, 7523, 7541, 7561, 7577, 7591, 7603, 7643, 7649, 7673, 7717, 7753, 7823, 7841, 7867, 7877, 7883, 7901, 7927, 7937, 8009, 8039, 8059, 8081, 8089, 8093, 8123, 8161, 8167, 8191, 8219, 8269, 8273, 8287, 8297, 8353, 8363, 8387, 8419, 8423, 8447, 8461, 8521, 8573, 8609, 8627, 8629, 8647, 8689, 8699, 8737, 8761, 8783, 8821, 8839, 8849, 8887, 8923, 8929, 8941, 8969, 9001, 9007, 9043, 9109, 9137, 9157, 9181, 9203, 9209, 9221, 9239, 9257, 9319, 9337, 9343, 9377, 9437, 9461, 9463, 9521, 9533, 9551, 9631, 9679, 9697, 9719, 9721, 9803, 9857, 9883, 9931, 10009, 10067, 10141, 10177, 10247, 10267, 10273, 10289, 10303, 10321, 10331, 10333, 10337, 10343, 10369, 10427, 10453, 10457, 10459, 10487, 10501, 10529, 10531, 10597, 10627, 10663, 10667, 10709, 10739, 10753, 10831, 10889, 10891, 10903, 10957, 10979, 10987, 11059, 11083, 11113, 11117, 11131, 11171, 11177, 11239, 11321, 11351, 11369, 11399, 11423, 11437, 11447, 11471, 11489, 11503, 11597, 11621, 11681, 11699, 11717, 11743, 11779, 11783, 11833, 11839, 11867, 11923, 11939, 11959, 11971, 11981, 12041, 12071, 12107, 12149, 12163, 12211, 12263, 12343, 12347, 12379, 12391, 12413, 12421, 12437, 12451, 12479, 12487, 12539, 12547, 12583, 12589, 12601, 12619, 12637, 12659, 12713, 12781, 12823, 12829, 12853, 12893, 12911, 12917, 12919, 12941, 13043, 13127, 13147, 13163, 13171, 13187, 13217, 13259, 13291, 13297, 13397, 13399, 13421, 13463, 13477, 13513, 13537, 13553, 13597, 13633, 13693, 13709, 13729, 13757, 13781, 13789, 13799, 13877, 13903, 13913, 13931, 13999, 14051, 14057, 14087, 14107, 14143, 14153, 14173, 14177, 14197, 14207, 14281, 14293, 14321, 14327, 14389, 14401, 14411, 14423, 14461, 14479, 14489, 14503, 14519, 14537, 14633, 14653, 14699, 14713, 14717, 14723, 14771, 14779, 14797, 14813, 14851, 14869, 14887, 14897, 14923, 14929, 15061, 15077, 15091, 15101, 15121, 15161, 15173, 15187, 15193, 15199, 15233, 15263, 15269, 15271, 15313, 15319, 15329, 15377, 15401, 15427, 15451, 15461, 15473, 15541, 15619, 15641, 15643, 15649, 15661, 15667, 15727, 15739, 15761, 15767, 15787, 15803, 15809, 15817, 15877, 15881, 15901, 15907, 15919, 15937, 15991, 16001, 16061, 16063, 16073, 16087, 16111, 16183, 16223, 16319, 16333, 16369, 16417, 16427, 16433, 16453, 16477, 16519, 16573, 16603, 16607, 16619, 16631, 16649, 16691, 16693, 16741, 16759, 16823, 16871, 16879, 16889, 16903, 16937, 16981, 16987, 17033, 17053, 17077, 17123, 17183, 17203, 17207, 17231, 17291, 17317, 17351, 17377, 17387, 17389, 17401, 17417, 17419, 17449, 17471, 17483, 17491, 17497, 17519, 17581, 17659, 17681, 17737, 17749, 17761, 17783, 17863, 17903, 17909, 17929, 17957, 17959, 17981, 18041, 18043, 18049, 18061, 18077, 18089, 18119, 18121, 18169, 18211, 18217, 18223, 18307, 18313, 18341, 18353, 18367, 18379, 18413, 18439, 18443, 18457, 18481, 18521, 18541, 18553, 18593, 18617, 18671, 18679, 18719, 18757, 18797, 18869, 18913, 18917, 18959, 19051, 19069, 19139, 19183, 19231, 19249, 19259, 19273, 19289, 19319, 19381, 19391, 19427, 19483, 19489, 19501, 19507, 19559, 19681, 19687, 19759, 19763, 19819, 19861, 19891, 19937, 19973, 20071, 20107, 20117, 20129, 20173, 20269, 20333, 20347, 20359, 20369, 20477, 20509, 20719, 20731, 20749, 20771, 20773, 20807, 20857, 20879, 20921, 20947, 20963, 20983, 21013, 21017, 21019, 21059, 21061, 21101, 21107, 21121, 21157, 21163, 21193, 21319, 21383, 21487, 21499, 21517, 21521, 21523, 21529, 21673, 21767, 21773, 21787, 21803, 21817, 21839, 21881, 21929, 21997, 22013, 22027, 22037, 22067, 22073, 22079, 22093, 22111, 22133, 22157, 22193, 22229, 22259, 22277, 22343, 22349, 22369, 22391, 22397, 22433, 22441, 22447, 22481, 22531, 22567, 22639, 22697, 22727, 22739, 22741, 22769, 22777, 22807, 22817, 22901, 22907, 22921, 22937, 22993, 23003, 23017, 23167, 23173, 23189, 23197, 23311, 23369, 23509, 23549, 23557, 23687, 23743, 23767, 23789, 23801, 23833, 23873, 23911, 23929, 23971, 23977, 24001, 24023, 24029, 24083, 24107, 24113, 24121, 24133, 24137, 24179, 24197, 24223, 24359, 24373, 24469, 24473, 24509, 24517, 24677, 24697, 24763, 24781, 24793, 24799, 24809, 24847, 24923, 24967, 24989, 25033, 25037, 25111, 25121, 25127, 25171, 25189, 25219, 25247, 25253, 25307, 25409, 25453, 25463, 25469, 25541, 25583, 25601, 25643, 25657, 25667, 25673, 25679, 25741, 25889, 25919, 25939, 25943, 25981, 25997, 26041, 26053, 26083, 26107, 26111, 26161, 26177, 26203, 26209, 26227, 26339, 26347, 26371, 26417, 26431, 26437, 26459, 26479, 26669, 26693, 26737, 26993, 27043, 27059, 27061, 27077, 27103, 27107, 27143, 27211, 27241, 27299, 27361, 27407, 27437, 27479, 27481, 27541, 27617, 27631, 27647, 27697, 27701, 27737, 27749, 27763, 27767, 27773, 27793, 27803, 27851, 27901, 27919, 27953, 27997, 28001, 28027, 28081, 28087, 28123, 28151, 28163, 28181, 28279, 28283, 28289, 28351, 28433, 28439, 28573, 28579, 28591, 28603, 28631, 28663, 28669, 28697, 28703, 28729, 28753, 28817, 28859, 28871, 28909, 28921, 28927, 28933, 28949, 29009, 29017, 29023, 29101, 29137, 29147, 29167, 29173, 29287, 29327, 29399, 29411, 29443, 29473, 29527, 29531, 29537, 29569, 29611, 29671, 29717, 29753, 29837, 29863, 29879, 29881}, 499500)))

}
