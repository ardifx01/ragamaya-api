package static

import (
	"ragamaya-api/api/predicts/dto"
	"strings"
)

type BatikPattern struct {
	Pattern     string `json:"pattern"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
	History     string `json:"history"`
}

var BatikPatterns = map[dto.Pattern]BatikPattern{
	dto.BALI: {
		Pattern:     "Bali",
		Origin:      "Bali",
		Description: "Batik Bali menonjolkan motif naturalistik dan dekoratif yang terinspirasi flora, fauna, dan ritual lokal; berkembang juga menjadi batik lukis dan cap berwarna cerah.",
		History:     "Industri batik Bali modern mulai tumbuh sekitar 1970-an dan salah satu pelopornya adalah Pande Ketut Krisna di Desa Batubulan, Gianyar. Pada fase awal pembuatannya lebih banyak menggunakan teknik cap dan alat tenun bukan mesin (ATBM); seiring waktu pengrajin Bali mengadaptasi teknik tulis, cap, dan lukis sehingga menghasilkan varian motif yang lebih berani secara koloristik dan komposisi — dari motif laut (ulamsari) sampai motif mitologis seperti Singa Barong. Perkembangan pariwisata mempercepat diversifikasi motif dan commercialisasi batik Bali sebagai produk budaya sekaligus komoditas ekonomi lokal; batik Bali hari ini menjadi kombinasi estetika tradisi dan desain kontemporer untuk pasar domestik dan internasional.",
	},
	dto.BETAWI: {
		Pattern:     "Betawi",
		Origin:      "Sunda Kelapa / Betawi (Greater Jakarta)",
		Description: "Batik Betawi menampilkan palet warna cerah dan motif yang merefleksikan identitas Betawi — ondel-ondel, pucuk rebung, tumpal — hasil akulturasi pesisir dan pengaruh Tionghoa/Arab.",
		History:     "Batik Betawi berkembang dari corak batik pesisir utara Jawa sejak abad ke-19 ketika Batavia menjadi titik pertemuan perdagangan dan budaya. Motif-motif khas Betawi terbentuk melalui asimilasi desain Jawa pesisir dengan elemen Tionghoa dan Timur Tengah, menjadikan batik Betawi berbeda dari batik keraton Jawa yang lebih kaku. Dalam praktik modern batik Betawi menjadi representasi identitas kota Jakarta dan mendapat perhatian dalam upaya pelestarian budaya lokal setelah pengakuan batik Indonesia sebagai warisan takbenda oleh UNESCO (2009). Perkembangan kontemporer menempatkan batik Betawi pada ranah fesyen kota dengan adaptasi motif untuk busana sehari-hari dan seragam institusi.",
	},
	dto.CELUP: {
		Pattern:     "Celup",
		Origin:      "Teknik: masuk lewat jalur perdagangan (Jumputan/ikat-celup dari Asia)",
		Description: "Celup merujuk pada teknik pewarnaan (jumputan/tie-dye/ikat celup), bukan satu motif; menghasilkan pola gradasi dan warna kuat melalui proses pencelupan.",
		History:     "Teknik celup/ikat-celup (jumputan) yang diaplikasikan untuk menghasilkan ragam motif masuk ke Nusantara melalui jalur perdagangan dari Tiongkok dan India dan diserap ke berbagai tradisi lokal. Di Indonesia teknik ini berkembang ke beragam nama dan tradisi (mis. jumputan di Jawa, sasirangan di Kalimantan, cinde di Sumatra). Celup menjadi alternatif produksi selain batik tulis dan cap karena lebih cepat dan bersifat efisien untuk produksi massal; pada masa kontemporer celupan dipakai baik untuk tradisi maupun kategori produk fashion yang lebih variatif.",
	},
	dto.CENDRAWASIH: {
		Pattern:     "Cendrawasih",
		Origin:      "Papua",
		Description: "Motif Cendrawasih menggambarkan burung cendrawasih yang endemik Papua—ikon identitas, keindahan, dan spiritualitas lokal.",
		History:     "Motif batik Cendrawasih lahir sebagai bentuk representasi fauna khas Papua yang kerap dianggap sakral dan simbol status spiritual/estetika komunitas lokal. Penggunaan gambar cendrawasih pada tekstil menguatkan narasi identitas regional—menghubungkan cerita lokal, ritual, dan sumber daya alam. Pada periode modern motif ini diadaptasi oleh perancang nusantara dan pengrajin untuk menyuarakan nilai pelestarian alam dan promosi budaya Papua dalam ranah nasional/internasional, tanpa mengabaikan konteks adat yang melekat pada ikon tersebut.",
	},
	dto.CEPLOK: {
		Pattern:     "Ceplok",
		Origin:      "Yogyakarta (pusat: Kotagede dan daerah Mataram)",
		Description: "Ceplok adalah kategori motif geometris berulang—unit-unit 'ceplok' (sekuntum/elemen berulang) berisi isian (isen-isen) yang membentuk pola simetris khas Jawa.",
		History:     "Motif Ceplok tergolong motif kuno yang berkembang sejak era kerajaan Mataram dan menjadi ciri batik klasik Yogyakarta. Bentuknya terdiri dari satuan berulang (lingkaran, bintang, segi empat) yang diisi dekorasi kecil (isen-isen) sehingga menciptakan permukaan yang harmonis dan simetris. Ceplok sering dipakai sebagai kain upacara dan pakaian resmi; variasinya (mis. Ceplok Kembang Kates) diproduksi di sentra Kayadangan/Kotagede. Dalam sejarahnya, ceplok berfungsi sebagai medium ekspresi estetika keraton sekaligus identitas masyarakat desa pembatik.",
	},
	dto.CIAMIS: {
		Pattern:     "Ciamis",
		Origin:      "Ciamis (Jawa Barat) — warisan budaya Priangan / Kerajaan Galuh)",
		Description: "Batik Ciamis menonjolkan motif bernuansa alam dan legenda lokal (Ciung Wanara, Galuh), biasanya dengan palet soga dan hitam pada dasarnya.",
		History:     "Tradisi batik Ciamis menautkan jejaknya pada Kerajaan Galuh dan praktik membatik yang berpindah/terbawa oleh migrasi sosial dan konflik (termasuk masa Perang Jawa). Intensitas produksi dan variasi motif mencapai puncaknya pada era 1960–1980; motif-motif seperti Ciung Wanara, Galuh Pakuan dan motif lokal lain merekam narasi sejarah setempat. Sejak akhir abad ke-20 penurunan jumlah perajin dan persaingan dari produksi cap/printing menimbulkan tantangan revitalisasi; namun studi etnografi dan program pelestarian mencatat nilai filosofis dan simbolik yang kuat pada batik Ciamis sehingga menjadi fokus program konservasi budaya.",
	},
	dto.GARUTAN: {
		Pattern:     "Garutan",
		Origin:      "Garut (Priangan Timur, Jawa Barat)",
		Description: "Batik Garutan khas Priangan Timur—menggabungkan motif flora/fauna dan ragam hias naturalistik dengan palet relatif landai.",
		History:     "Batik Garutan adalah tradisi panjang masyarakat Garut yang mendapat momentum produksi lebih besar pada abad ke-19 sampai abad ke-20; pada masa kolonial tokoh seperti Karel F. Holle memberi perhatian terhadap kerajinan lokal sehingga industri kecil berkembang. Batik Garutan mencapai masa jayanya antara 1967–1985; motif seperti Merak Ngibing dan ragam hias khas Garut memantapkan identitasnya. Krisis bahan, perubahan pasar, dan modernisasi mengakibatkan penurunan produksi; belakangan muncul gerakan revitalisasi berbasis warisan lokal dan branding daerah untuk mengembalikan posisi Garut dalam peta batik Priangan.",
	},
	dto.GENTONGAN: {
		Pattern:     "Gentongan",
		Origin:      "Tanjung Bumi, Bangkalan (Madura)",
		Description: "Batik Gentongan identik dengan Madura—motifnya sederhana hingga naturalistik, dan istilah 'gentongan' berkaitan dengan wadah (gentong) yang dipakai dalam proses pewarnaan.",
		History:     "Batik Gentongan berkembang di komunitas pesisir Madura (Tanjung Bumi, Bangkalan) dan terkait dengan teknik pewarnaan tradisional yang menggunakan gentong untuk merendam/menyimpan cairan pewarna. Sejarah lisan dan tulisan menunjukkan produksi bertumbuh pada masa awal abad ke-20 dan motif khas seperti Tong Centong (alat penyendok nasi) lahir di pertengahan abad ke-20 (sekitar 1950-an) sebagai tanda fungsi sosial (pemberian hantaran pernikahan). Ragam gentongan dipengaruhi gaya hidup pelaut/pedagang dan kerap memakai warna cerah yang menjadi ciri pengenal Madura.",
	},
	dto.KAWUNG: {
		Pattern:     "Kawung",
		Origin:      "Yogyakarta / Kerajaan Mataram (Jawa Tengah/Yogyakarta)",
		Description: "Kawung adalah motif geometris tertua yang berulang menyerupai buah kawung/kolang-kaling atau teratai; dipandang sebagai simbol kemurnian, keseimbangan, dan kedaulatan keraton.",
		History:     "Motif Kawung diperkirakan berasal dari tradisi kuno Jawa dan dikaitkan dengan lingkup kekuasaan Mataram dan keraton; beberapa sumber menyebut keterkaitan dengan era Majapahit/Mataram sehingga dianggap salah satu motif tertua. Secara historis Kawung dipakai pada lingkungan keraton sebagai tanda status dan makna etis—kesempurnaan dan tata krama. Seiring waktu variasi kawung (picis, bribil, semar dsb.) muncul dan motif ini menjadi salah satu pola paling dihormati dalam kanon batik klasik Jawa.",
	},
	dto.KERATON: {
		Pattern:     "Keraton",
		Origin:      "Lingkungan Keraton (Surakarta & Yogyakarta)",
		Description: "Batik Keraton (vorstenlanden) adalah kelompok motif dan aturan pakai yang lahir di lingkungan istana — sarat simbolisme, palet sogan/indigo, dan kaidah pakem.",
		History:     "Batik keraton berkembang sejak masa Mataram Islam (abad ke-16–17) dan menjadi pusat kodifikasi estetika tekstil Jawa: motif, warna, dan tata pemakaian diatur ketat oleh norma istana. Motif-motif keraton seperti Parang, Kawung, Sidomukti, Sidoluhur memiliki peruntukan ritual dan sosial (mis. pakaian resmi, upacara, penanda status). Proses translasi dari kain keraton ke masyarakat umum terjadi secara bertahap—mulai pelunakan larangan hingga adaptasi untuk konsumsi luas pada abad ke-19–20—tetapi nilai simbolik keraton tetap menjadi rujukan utama dalam kajian batik klasik.",
	},
	dto.LASEM: {
		Pattern:     "Lasem",
		Origin:      "Lasem, Pesisir Utara Jawa (Rembang area, Jawa Tengah)",
		Description: "Batik Lasem dikenal sebagai hasil akulturasi Jawa–Tionghoa: palet merah khas, ornamen naga/hong, dan estetika pesisir.",
		History:     "Lasem menjadi titik pertemuan budaya karena komunitas Tionghoa yang menetap sejak lama; pengaruh ini tercermin pada motif dan warna batik Lasem (merah terang, elemen oriental seperti naga dan pipi ceramic). Sejarah perkembangannya menunjukkan keterkaitan kepengaruhan Rembang-Lasem, migrasi Tionghoa, dan kontak perdagangan sejak abad ke-13–18; keahlian pembuatan batik Lasem kemudian bertransformasi menjadi produk unggulan yang diekspor dan dikonsumsi lintas komunitas. Di era modern Lasem tetap dikenal sebagai sentra batik pesisir dengan ciri khas ornamentasi oriental yang kuat.",
	},
	dto.MEGAMENDUNG: {
		Pattern:     "Megamendung",
		Origin:      "Cirebon (Pesisir Utara Jawa Barat)",
		Description: "Motif awan berlapis—megamendung—adalah ikon batik Cirebon yang memvisualisasikan awan, hujan, dan pengaruh estetik Tionghoa.",
		History:     "Megamendung lahir di Cirebon sebagai produk akulturasi visual antara motif awan Tionghoa dan tradisi batik lokal—pengaruh ini diperkuat oleh hubungan kerajaan Cirebon dengan komunitas Tionghoa (mis. pernikahan Sunan Gunung Jati dengan Ratu Ong Tien dalam tradisi lokal). Motif ini awalnya dipakai dalam kontekstual kerajaan/pesisir dan kemudian menyebar ke produksi rakyat; filosofi megamendung memuat makna kesabaran, harapan akan hujan (kesuburan), dan harmoni alam-manusia. Kini megamendung menjadi simbol identitas Cirebon dan sering diadaptasi dalam desain kontemporer.",
	},
	dto.PARANG: {
		Pattern:     "Parang",
		Origin:      "Kartasura / Kerajaan Mataram (Solo/Yogyakarta)",
		Description: "Parang adalah motif diagonal bergelombang (mirip bilah parang/lereng) yang melambangkan kekuatan, kesinambungan, dan etika perjuangan.",
		History:     "Motif Parang dikaitkan kuat dengan Sultan Agung (kerajaan Mataram) dan tradisi keraton pada abad ke-17; variasi seperti Parang Rusak, Parang Barong, Klithik memiliki fungsi simbolik berbeda serta peruntukan status sosial. Parang pada mulanya dibatasi penggunaannya di lingkungan istana (motif larangan) dan dipakai untuk menyampaikan pesan moral: keteguhan, pengendalian diri, dan kesinambungan generasi. Dalam perkembangan masyarakat Parang menyebar ke kalangan bangsawan dan rakyat serta menjadi salah satu motif klasik paling berpengaruh dalam kanon batik Jawa.",
	},
	dto.PEKALONGAN: {
		Pattern:     "Pekalongan",
		Origin:      "Pekalongan (Jawa Tengah) — batik pesisir 'Tujuh Rupa'",
		Description: "Batik Pekalongan (termasuk Tujuh Rupa) kaya warna dan motif flora/fauna; kuat dipengaruhi kontak perdagangan dan multikultural pelabuhan.",
		History:     "Sebagai pelabuhan yang menerima aliran pedagang asing, Pekalongan sejak abad ke-19 menjadi pusat batik pesisir yang mengakomodasi elemen Tionghoa, Arab, India, dan Eropa. Motif Tujuh Rupa (tujuh rupa) menonjolkan komposisi flora/fauna dan hiasan keramik Tionghoa yang diadaptasi secara lokal; teknik produksi berkembang (colet/kuas, cap, printing) untuk memenuhi permintaan massal. Sejak abad ke-20 industri batik Pekalongan mengalami industrialisasi kuantitas sekaligus inovasi motif, menjadikannya salah satu sentra batik terbesar di Indonesia.",
	},
	dto.PRIANGAN: {
		Pattern:     "Priangan",
		Origin:      "Priangan / Parahyangan (Jawa Barat: Tasikmalaya, Garut, Ciamis area)",
		Description: "Batik Priangan (Sundanese) memadukan estetika Sunda—motif naturalistik dan palet tertentu—dengan sejarah lokal Tarumanegara dan pengaruh Mataram.",
		History:     "Jejak batik di Priangan diperkirakan sudah ada jauh sebelum era kolonial; motif Priangan berkembang melalui kontak internal (migrasi pengrajin dari Jawa Tengah pada masa perang, mis. awal abad ke-19) dan warisan kerajaan lokal. Produksi batik Priangan sempat naik turun (pasang-surut akibat perubahan ekonomi dan kompetisi kain printing), namun beberapa motif khas—termasuk Merak Ngibing—menguat sebagai identitas lokal. Revitalisasi kontemporer menekankan praktik pelatihan, branding lokal, dan integrasi desain modern untuk menjaga kesinambungan industri kerajinan Priangan.",
	},
	dto.SEKAR: {
		Pattern:     "Sekar",
		Origin:      "Keraton (Solo/Yogyakarta) — dikenal sebagai 'Sekar Jagad'",
		Description: "Sekar (Sekar Jagad) adalah ragam motif bunga dunia—patch-like floral composition yang melambangkan keragaman dan keindahan dunia.",
		History:     "Sekar Jagad muncul dan populer di lingkungan keraton sejak abad ke-18; motif ini bersifat tidak teratur (island-like lines) dan sering memuat isen-isen yang beragam (kawung, truntum, flora/fauna) sehingga menyerupai peta dunia bunga. Fungsi historisnya bersifat representasional: simbol keberagaman, harmoni, dan ekspresi estetika yang lebih bebas dibanding motif keraton lain yang sangat pakem. Sekar tetap diaplikasikan dalam tekstil upacara dan menjadi sumber inspirasi desain kontemporer karena komposisi ornamentalnya yang kaya.",
	},
	dto.SIDOLUHUR: {
		Pattern:     "Sidoluhur",
		Origin:      "Keraton Surakarta / Yogyakarta (lazim pada tradisi keraton Jawa)",
		Description: "Sidoluhur adalah motif keraton yang identik dengan kehormatan, teladan, dan biasanya dipakai pada acara-acara resmi atau pernikahan.",
		History:     "Sidoluhur merupakan bagian dari kanon motif keraton yang diwariskan dalam lingkungan istana Kartasura/Surakarta dan Yogyakarta sejak periode Mataram; motif ini sering dianggap membawa pesan moral—agar pemakai menjadi panutan yang luhur. Dalam praktik upacara Sidoluhur dipilih untuk pakaian resmi (termasuk perlengkapan prosesi perkawinan) dan disertai ritual penggunaan khusus. Literatur etnografi keraton merekam Sidoluhur sebagai salah satu motif dengan makna sosial yang kuat dalam tata upacara Jawa.",
	},
	dto.SIDOMUKTI: {
		Pattern:     "Sidomukti",
		Origin:      "Keraton Surakarta / Yogyakarta",
		Description: "Sidomukti (dari 'sido' + 'mukti') bermakna harapan kemakmuran dan kehormatan — sering dipakai dalam konteks pernikahan dan upacara keluarga.",
		History:     "Sidomukti adalah motif turun-temurun di lingkungan keraton Surakarta/Yogyakarta, dikembangkan sebagai variasi dari motif keraton klasik (terkadang disebut turunan Sidomulyo). Motif ini tradisionalnya digunakan dalam upacara perkawinan sebagai lambang doa untuk kesejahteraan keluarga dan status sosial yang terjaga. Seiring modernisasi, penggunaan Sidomukti tetap bertahan pada prosesi adat namun juga diadaptasi untuk produk komersial dengan tetap menjaga elemen simboliknya.",
	},
	dto.SOGAN: {
		Pattern:     "Sogan",
		Origin:      "Yogyakarta / Solo (warna dasar 'soga' dari pewarna kayu)",
		Description: "Sogan merujuk pada palet warna cokelat-kekuningan yang dihasilkan dari pewarna alami (kayu soga); lazim di batik keraton Jawa.",
		History:     "Batik Sogan adalah salah satu tipe klasik yang khas di lingkungan keraton—ciri utamanya adalah penggunaan warna soga (cokelat keemasan) yang berasal dari ekstrak batang kayu tertentu. Secara simbolik warna-warna sogan berhubungan dengan kesakralan, kebijakan etika, dan kesederhanaan Jawa yang berorientasi batin; batik ini banyak digunakan dalam pakaian upacara keraton. Di era modern batik sogan dipertahankan sebagai identitas visual batik klasik, sekaligus menjadi bahan studi pewarnaan alami dan konservasi teknik tradisional.",
	},
	dto.TAMBAL: {
		Pattern:     "Tambal",
		Origin:      "Yogyakarta / Jawa Tengah (jenis komposit/pola 'patch')",
		Description: "Tambal secara harfiah berarti 'patch' — batik tambal adalah komposisi berbagai motif yang disusun seperti tambalan atau patchwork.",
		History:     "Tambal berkembang sebagai teknik komposit yang menggabungkan berbagai potongan kain atau motif menjadi satu kesatuan; dalam tradisi, kain tambal pernah dipakai sebagai kain penutup bagi orang sakit atau untuk tujuan ritual penyembuhan karena dipercaya membawa energi penyembuhan. Motif Tambal (termasuk varian Tambal Seribu) mencerminkan filosofi perbaikan dan kontinuitas—manusia harus selalu berusaha memperbaiki diri. Seiring ketahanan budaya, tambal dipertahankan sebagai gaya dekoratif dalam batik kontemporer dan dipakai pula untuk proyek revitalisasi kultural.",
	},
}

func GetBatikPattern(name dto.Pattern) (*BatikPattern, bool) {
	if pattern, exists := BatikPatterns[name]; exists {
		return &pattern, true
	}

	return nil, false
}

func GetAllBatikPatterns() []BatikPattern {
	patterns := make([]BatikPattern, 0, len(BatikPatterns))
	for _, pattern := range BatikPatterns {
		patterns = append(patterns, pattern)
	}
	return patterns
}

func SearchByOrigin(origin string) []BatikPattern {
	var results []BatikPattern
	lowerOrigin := strings.ToLower(origin)

	for _, pattern := range BatikPatterns {
		if strings.Contains(strings.ToLower(pattern.Origin), lowerOrigin) {
			results = append(results, pattern)
		}
	}
	return results
}
