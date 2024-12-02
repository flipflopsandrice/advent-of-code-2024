package one

import (
	"testing"
)

func TestDistanceSum(t *testing.T) {
	sum := distanceSum(
		[]int{19394, 88523, 60736, 75224, 78157, 40925, 86284, 28269, 95540, 48998, 30875, 39762, 90841, 72017, 90563, 97482, 41285, 10561, 55289, 58094, 25799, 39530, 26398, 84654, 17891, 49607, 77920, 63548, 49452, 10623, 44369, 24171, 23836, 35024, 53006, 17908, 96110, 21260, 19069, 50462, 58898, 63205, 68401, 13938, 97309, 79019, 80069, 62881, 93649, 62227, 67133, 22177, 70120, 65463, 20788, 16101, 88771, 17935, 39580, 13624, 39315, 87396, 39438, 19555, 12294, 39136, 68137, 64762, 11693, 27873, 74977, 65309, 99968, 73977, 38300, 69852, 65284, 43782, 74032, 37820, 39747, 85435, 13406, 49130, 52354, 92149, 36943, 12471, 60136, 44887, 87542, 46184, 49717, 96962, 34123, 51911, 29657, 78653, 59461, 29224, 36439, 13523, 90606, 25442, 81801, 23207, 47997, 98666, 79422, 93922, 55628, 89976, 86134, 66713, 70140, 44831, 94934, 51449, 69165, 69254, 23289, 25055, 14060, 55286, 37866, 41143, 82404, 64498, 19413, 91738, 98188, 18533, 26451, 27689, 63800, 51092, 33392, 61939, 93244, 69356, 30830, 73378, 54810, 60906, 64786, 73954, 51826, 45477, 39832, 39313, 87150, 24195, 60222, 53575, 12726, 70517, 89448, 94123, 33021, 12660, 38086, 47154, 62889, 26260, 62539, 10837, 32785, 61750, 48843, 23001, 78135, 46469, 92523, 77332, 88022, 69886, 76828, 30411, 37813, 58166, 11224, 47464, 44093, 18663, 67651, 82915, 84939, 93163, 35337, 31836, 83607, 27234, 17611, 45939, 44136, 24980, 64751, 82530, 83164, 52225, 39901, 68629, 41761, 84332, 56294, 88603, 34147, 70748, 83432, 61677, 55026, 45757, 53241, 31385, 56659, 60236, 65487, 57628, 38867, 61389, 52190, 73770, 28535, 72684, 68234, 70718, 57893, 87278, 99048, 85431, 21498, 11733, 29786, 83685, 84668, 96742, 83192, 14291, 15967, 40857, 37360, 28437, 81625, 10670, 39088, 79068, 86163, 16356, 77035, 95893, 68524, 67206, 61332, 18337, 93390, 54333, 71139, 38055, 68004, 77905, 47611, 79842, 13662, 53745, 41222, 52076, 15371, 23306, 29817, 93289, 22128, 38431, 23505, 28107, 86820, 26522, 50144, 82874, 66228, 80001, 42537, 39967, 85939, 65039, 20791, 13246, 29704, 37362, 52727, 56123, 31682, 99721, 11915, 78695, 37842, 33522, 42425, 61796, 18581, 85444, 79723, 32198, 61485, 25570, 59494, 31201, 20105, 68086, 99087, 69972, 62874, 76958, 15510, 32628, 98849, 57502, 99195, 71072, 70180, 90411, 17690, 72498, 62472, 23478, 21680, 64962, 67171, 87749, 96550, 34827, 14965, 81400, 28022, 65594, 61051, 37268, 42755, 90944, 14515, 78631, 83844, 47511, 63833, 27770, 63009, 58607, 72299, 98414, 59427, 42920, 42943, 76950, 66370, 62072, 69498, 38482, 90604, 79371, 24310, 84526, 28190, 42959, 62789, 96930, 17482, 36117, 87059, 21860, 66292, 72327, 32333, 95478, 40395, 23169, 46662, 27483, 14770, 30201, 81414, 71539, 18733, 71233, 23912, 72646, 80943, 53771, 25524, 71583, 68782, 92398, 18717, 71309, 66027, 92188, 52654, 41484, 74536, 34402, 65396, 64587, 73313, 64781, 40433, 15383, 53862, 18251, 92391, 42843, 30265, 70472, 11115, 38743, 29204, 89512, 47347, 35918, 35111, 87122, 62610, 69117, 34608, 68159, 44645, 11360, 98841, 68166, 30855, 34886, 19259, 49446, 28770, 69727, 19805, 13092, 21397, 79152, 68649, 95088, 38088, 34497, 36209, 30822, 23427, 18266, 42203, 86082, 31027, 41998, 15099, 11613, 59774, 38313, 73717, 98902, 99167, 41435, 15042, 78981, 76225, 67283, 84631, 61343, 12940, 57352, 73255, 65000, 79110, 55163, 36832, 92421, 24862, 13598, 76781, 64986, 48307, 30499, 18397, 44296, 53975, 46262, 28139, 49533, 93528, 13765, 72839, 90780, 69774, 45153, 83124, 91418, 70779, 78588, 96799, 17454, 39594, 96631, 62254, 32818, 63705, 78591, 39160, 75105, 91888, 64917, 79697, 23416, 69209, 86468, 87174, 21197, 83503, 46243, 83425, 64091, 82387, 72432, 68423, 63036, 46945, 92178, 92475, 71265, 22466, 78425, 84989, 59915, 62251, 25112, 78278, 27294, 49019, 21504, 11920, 33923, 45471, 59325, 80028, 25706, 81065, 14435, 20330, 51339, 72276, 83903, 80068, 66905, 79932, 19408, 83911, 84717, 35764, 91958, 12711, 61725, 39768, 86115, 17195, 99304, 95620, 23299, 70384, 95974, 19983, 30328, 93287, 29850, 71063, 87005, 13402, 16810, 53126, 87495, 59241, 38729, 97527, 54481, 53294, 99814, 52458, 97173, 90268, 15958, 27539, 75668, 92730, 54793, 20493, 38749, 63488, 69715, 99423, 13503, 37032, 82994, 95113, 82341, 47782, 19392, 58145, 11367, 24090, 74367, 58821, 84758, 41324, 63755, 18456, 43908, 76789, 13654, 92979, 45134, 39852, 55066, 53972, 35205, 34894, 93606, 45212, 71473, 77246, 72899, 39506, 13959, 56571, 68789, 29980, 82077, 95915, 31832, 79076, 47570, 21058, 19787, 36452, 24331, 35181, 65847, 28677, 30918, 71808, 78490, 50658, 84525, 92052, 35798, 67640, 28644, 79569, 55849, 60737, 78807, 49986, 61356, 85575, 43531, 53843, 92048, 18275, 59311, 63750, 77779, 26196, 65202, 54245, 60651, 12795, 12691, 14868, 57009, 22924, 46880, 44046, 88721, 62559, 16394, 82812, 91964, 24111, 54143, 96970, 65895, 90345, 11847, 26002, 15535, 33281, 54737, 75392, 43098, 41574, 87423, 76419, 65498, 10938, 45987, 54998, 37518, 23092, 56921, 84101, 23120, 53441, 10525, 43726, 61391, 68630, 62125, 53111, 39402, 51205, 58826, 99060, 35880, 92110, 65717, 30095, 51486, 67902, 18402, 29938, 18611, 85078, 64358, 92243, 22205, 78191, 66220, 59143, 37247, 66923, 77849, 78619, 15016, 51214, 60173, 40398, 85099, 14582, 27504, 99400, 11061, 71600, 65520, 46348, 82323, 60500, 50789, 29819, 53629, 68689, 66009, 66056, 79406, 16799, 81044, 51166, 82919, 14096, 22876, 74460, 35827, 26421, 26223, 41827, 16656, 82262, 20908, 35562, 57341, 76022, 79301, 53092, 92740, 52415, 23043, 68913, 25155, 89907, 57055, 83130, 62554, 78177, 42006, 59006, 54514, 91441, 96678, 85071, 65555, 17750, 95125, 79124, 91205, 66332, 65100, 65055, 45502, 53781, 12902, 89582, 42262, 58411, 45418, 24720, 59590, 96877, 53645, 41081, 44928, 14232, 88042, 48259, 80858, 22758, 76760, 75606, 38843, 79989, 97934, 11337, 94685, 11546, 44205, 21374, 78416, 34842, 22383, 80155, 93228, 88857, 84048, 19980, 89184, 35848, 17400, 10985, 45857, 69026, 24044, 68840, 18020, 12012, 53078, 22616, 83684, 69679, 18547, 61762, 32214, 54647, 49923, 30699, 82213, 89544, 86907, 10007, 33723, 71326, 43056, 22112, 29413, 62103, 99839, 82466, 63961, 53925, 76255, 63941, 99403, 16280, 24735, 80271, 54165, 89734, 14857, 75370, 97725, 48788, 76858, 31339, 74171, 92778, 81891, 32207, 89414, 46799, 31925, 30646, 93120, 12245, 57323, 12723, 87211, 87213, 24725, 48091, 90573, 33839, 84427, 76722, 33739, 26953, 18734, 50146, 74422, 11801, 99794, 43452, 88161, 63363, 85697, 20311, 87978, 70962, 95376, 39600, 49544, 86928, 36128, 26726, 80767, 98537, 96829, 90261, 22485, 62557, 90304, 49968, 41017, 89768, 56055, 70937, 37132, 86318, 51330, 74069, 35200, 14871, 88542, 70691, 52454, 96898, 13342, 51301, 31504, 89361, 66758, 76024, 33179, 90757, 90820, 10606, 38801, 20602, 73304, 15768, 42580, 89093, 42484, 75272, 56843, 33767, 43275, 95051, 40784, 92406, 24412, 75949, 23180, 95260, 87848, 37421, 12916, 52570, 63243, 41946, 15841, 32242, 84236, 63032, 25207, 30525, 38044, 17774, 25953, 31123, 88196, 13796, 19398, 88811, 71454, 94597, 14793, 86183, 49244, 11382, 33929, 15981, 53640},
		[]int{30201, 40612, 82993, 92188, 27314, 54647, 23731, 14815, 16406, 59590, 22128, 70353, 51911, 77153, 12816, 68004, 32178, 31048, 51911, 14440, 18533, 88299, 56543, 62539, 55936, 98934, 33887, 68649, 45428, 68630, 79502, 42203, 54647, 97148, 14515, 12813, 45282, 31585, 68649, 59959, 80858, 27250, 14973, 96725, 16113, 22128, 51105, 31929, 28456, 87150, 38055, 28210, 41359, 22128, 53022, 93317, 80239, 96253, 59590, 14515, 81433, 39312, 36516, 37518, 73072, 39440, 38055, 54647, 39768, 21222, 22128, 37518, 93649, 51911, 69694, 18529, 42203, 30201, 80207, 68004, 42203, 18611, 52826, 98817, 55049, 43275, 93096, 50411, 24044, 12366, 84432, 30378, 55951, 43275, 51911, 22128, 16830, 14100, 38055, 38218, 68630, 39762, 38055, 34772, 16055, 10144, 99304, 78425, 62539, 16723, 61077, 79887, 46710, 69568, 97593, 34608, 44979, 36022, 18010, 68649, 15723, 39762, 43275, 15464, 38055, 36251, 43275, 10231, 77772, 36203, 68649, 22128, 14515, 18533, 28677, 84147, 95229, 29850, 63867, 44901, 68630, 22118, 98128, 89093, 49722, 68004, 64510, 17137, 19305, 77094, 62539, 19958, 55319, 18533, 80970, 95031, 52996, 17400, 58607, 49866, 13197, 55140, 68004, 79122, 27869, 68630, 80193, 83026, 16021, 86838, 38055, 70018, 24044, 48312, 59590, 63430, 99616, 68630, 44203, 68004, 90390, 89093, 68004, 94925, 92188, 68004, 72354, 47611, 80589, 41400, 54647, 28677, 17933, 42203, 88447, 76214, 83636, 59590, 96240, 71708, 56044, 54625, 39768, 39762, 22128, 30201, 80858, 54647, 42203, 53864, 29850, 39762, 93649, 37518, 92510, 13326, 22847, 29627, 28677, 81881, 28677, 21680, 28677, 93649, 51911, 61655, 14515, 88350, 38055, 14513, 38055, 75489, 14018, 28535, 28677, 46929, 93649, 80858, 70441, 33823, 24044, 87500, 87150, 25629, 59590, 70485, 92549, 57690, 34542, 30201, 12245, 44145, 70163, 20493, 23528, 28535, 68630, 68004, 38055, 75926, 25005, 25140, 24044, 68004, 51558, 81725, 68630, 93649, 62872, 39762, 20318, 89215, 14273, 58177, 68059, 92188, 92577, 12814, 43275, 51911, 38055, 43275, 62539, 58483, 26022, 71961, 51911, 28677, 46073, 51911, 80858, 80046, 22128, 37518, 63797, 18533, 66307, 22896, 19401, 99628, 44667, 99245, 68649, 81282, 49927, 28027, 62342, 24044, 94172, 57265, 85761, 86675, 32803, 75816, 14515, 74495, 25898, 55730, 54871, 60473, 53384, 20928, 37518, 59585, 17400, 92188, 38817, 88305, 31707, 37518, 28677, 37518, 68630, 84915, 14515, 22708, 36626, 68004, 62539, 59726, 21457, 35413, 74936, 46589, 62539, 32317, 58657, 68649, 46701, 37518, 34608, 81768, 95413, 68649, 21093, 80380, 89093, 93649, 29039, 43555, 77592, 25808, 65719, 62539, 83419, 13993, 68630, 84317, 18533, 54647, 80858, 92598, 20333, 68649, 34608, 49185, 28677, 54647, 33531, 68004, 97799, 72190, 11930, 59590, 24568, 34976, 83036, 94321, 83197, 93066, 34608, 17400, 21680, 20493, 18533, 57624, 39762, 41766, 46662, 54185, 20493, 72433, 51135, 80240, 92188, 57427, 61937, 92188, 10605, 91659, 38423, 68649, 27470, 14515, 37341, 91685, 12245, 12864, 18533, 39762, 90733, 47235, 62103, 55339, 54077, 93649, 40539, 53682, 20845, 62103, 85627, 46662, 35760, 22128, 22128, 54647, 10146, 94375, 53889, 57303, 35946, 14481, 65318, 93649, 37288, 73722, 51911, 92188, 14515, 25759, 40495, 39768, 99304, 40528, 10317, 61425, 67261, 78425, 83595, 22175, 93649, 28178, 24044, 92188, 82693, 64876, 12189, 93265, 99304, 78808, 22772, 50315, 67526, 89093, 14515, 81298, 72464, 37518, 18611, 24044, 37518, 38055, 24044, 92179, 79453, 21608, 54647, 46074, 17400, 40532, 80858, 78914, 87616, 65856, 42203, 21680, 28677, 42203, 43275, 78425, 78425, 63469, 28677, 39762, 93649, 73388, 38517, 15028, 93649, 42042, 62539, 64759, 42203, 59590, 38055, 61526, 44628, 34059, 69634, 23198, 24044, 17669, 93938, 33739, 67532, 24044, 76757, 18533, 78229, 94970, 56614, 43275, 20613, 78425, 46569, 93446, 18533, 18482, 54169, 59686, 39762, 31931, 68004, 38055, 14515, 42203, 57636, 21680, 99103, 46260, 31880, 38055, 89093, 95957, 29557, 18533, 99304, 46662, 75784, 51888, 94672, 37909, 64539, 43275, 67609, 93226, 88634, 58583, 28677, 20493, 43275, 53455, 22128, 36581, 22250, 87922, 41392, 43405, 64149, 28535, 65954, 96885, 94140, 53320, 76721, 64662, 43955, 26031, 67792, 74163, 25406, 92188, 74091, 87542, 23287, 14082, 68649, 68109, 34558, 37518, 93649, 51289, 67880, 49564, 43221, 12136, 92213, 15825, 99304, 28677, 67378, 55863, 28677, 47477, 28677, 57113, 37518, 39762, 68223, 76746, 49748, 39768, 30201, 92188, 96952, 89676, 21680, 30201, 49831, 97611, 92188, 12245, 62539, 19635, 37518, 83335, 30201, 93322, 38055, 21680, 47589, 89093, 46493, 30201, 28677, 29805, 78425, 68649, 28677, 64863, 54647, 49913, 38055, 68004, 14515, 86349, 22835, 98829, 42231, 59590, 93503, 51911, 49215, 22128, 71063, 54736, 24044, 52335, 76240, 93649, 86143, 83204, 54647, 81171, 60554, 14138, 51911, 89572, 75105, 39762, 95647, 92157, 79560, 78425, 32659, 44103, 68649, 43275, 93649, 24044, 29446, 22128, 75360, 76890, 28677, 14515, 26279, 26076, 89093, 60094, 18301, 99304, 80129, 56782, 40686, 24829, 91104, 43275, 90283, 16948, 34608, 61814, 75190, 98594, 50013, 52474, 63994, 39762, 84259, 58434, 17400, 25182, 33378, 32171, 68919, 16964, 75823, 17400, 38586, 30201, 89998, 99304, 18533, 78465, 51911, 38229, 54647, 18533, 55523, 42203, 19531, 76191, 63853, 78425, 39762, 68649, 13832, 93649, 62539, 12245, 50868, 71159, 11686, 17400, 28677, 41209, 53906, 22523, 17400, 98879, 81300, 26565, 68630, 38055, 68004, 51911, 24044, 18533, 14052, 52417, 86260, 64857, 21680, 22128, 28981, 20493, 32022, 38763, 71063, 59590, 17400, 86934, 90086, 29151, 59590, 62539, 78587, 33768, 73073, 24044, 68004, 66725, 93495, 18533, 73755, 68630, 38055, 68630, 39762, 98857, 53003, 68004, 94659, 68392, 59590, 73476, 15344, 38315, 28871, 29773, 68649, 92094, 46083, 93649, 88432, 18611, 54804, 78425, 62539, 54647, 46662, 23016, 83071, 30003, 38055, 17400, 39201, 51911, 43275, 53319, 22128, 33319, 29201, 21680, 84371, 28677, 45969, 54403, 57795, 24044, 18719, 46662, 94899, 99304, 57170, 80858, 37518, 21689, 48895, 40187, 37518, 68630, 84626, 43275, 40557, 18533, 92188, 31433, 59590, 23659, 93649, 34871, 28890, 72605, 34608, 98699, 37518, 87718, 54647, 31572, 22238, 29508, 12245, 90392, 58572, 29850, 24044, 37518, 48744, 68217, 49004, 92188, 62539, 46662, 42203, 17400, 43275, 15424, 80858, 60565, 87632, 87542, 44094, 15455, 22128, 87661, 52235, 62167, 25755, 22128, 37892, 56210, 28610, 71063, 93649, 75490, 80858, 69569, 17400, 18611, 72507, 20493, 89093, 44496, 51911, 38555, 38055, 80007, 25847, 89093, 68649, 87542, 34966, 26089, 62539, 92188, 67857, 51894, 82644, 54386, 30201, 80995, 92185, 68459, 59840, 18167, 24044, 34718, 17547, 39762, 68630, 54647, 17400, 83750, 51440, 14515, 77628, 99304, 68630, 20780, 15919, 85981, 92699, 79056, 91576, 39768, 42203, 45606, 37081, 68649, 21554, 68649, 56156, 43275, 85058, 22128, 30201, 99304, 16311, 39762, 45202, 87542, 93075, 39425, 49541, 92786, 18533, 64834, 18611, 87067, 68649, 93649, 37518, 17787, 33739, 20493, 43275, 42203, 77501, 80858, 11201, 62539, 52436, 16357, 77981, 92146, 76204, 51911, 68004, 45703},
	)

	if sum != 2742123 {
		t.Errorf("Expected 11, got %d", sum)
	}
}
