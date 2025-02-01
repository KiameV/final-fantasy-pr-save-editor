package ff4

import (
	"pixel-remastered-save-editor/models"
)

var (
	Bestiary = []models.NameValue{models.NewValueName(1, "Goblin"),
		models.NewValueName(2, "Basilisk"),
		models.NewValueName(3, "Helldiver"),
		models.NewValueName(4, "Floating Eyeball"),
		models.NewValueName(5, "Insectus"),
		models.NewValueName(6, "Cave Bat"),
		models.NewValueName(7, "Treant"),
		models.NewValueName(8, "Sword Rat"),
		models.NewValueName(9, "Tiny Mage"),
		models.NewValueName(10, "Desert Sahagin"),
		models.NewValueName(11, "Flying Eyes"),
		models.NewValueName(12, "Hundlegs"),
		models.NewValueName(13, "Sand Worm"),
		models.NewValueName(14, "Gigantoad"),
		models.NewValueName(15, "Vile Shell"),
		models.NewValueName(16, "Zombie"),
		models.NewValueName(17, "Killer Fish"),
		models.NewValueName(18, "Blood Flower"),
		models.NewValueName(19, "Toadgre"),
		models.NewValueName(20, "Red Mousse"),
		models.NewValueName(21, "Yellow Jelly"),
		models.NewValueName(22, "Sahagin"),
		models.NewValueName(23, "Leshy"),
		models.NewValueName(24, "Chrysalis"),
		models.NewValueName(25, "Goblin Captain"),
		models.NewValueName(26, "Black Lizard"),
		models.NewValueName(27, "Adamantoise"),
		models.NewValueName(28, "Domovoi"),
		models.NewValueName(29, "Purple Bavarois"),
		models.NewValueName(30, "Baron Soldier"),
		models.NewValueName(31, "General"),
		models.NewValueName(32, "Gatlinger"),
		models.NewValueName(33, "Li'l Murderer"),
		models.NewValueName(34, "Death Flower"),
		models.NewValueName(35, "Spirit"),
		models.NewValueName(36, "Skeleton"),
		models.NewValueName(37, "Cockatrice"),
		models.NewValueName(38, "Gargoyle"),
		models.NewValueName(39, "Roc Baby"),
		models.NewValueName(40, "Belphegor"),
		models.NewValueName(41, "Zu"),
		models.NewValueName(42, "Water Bug"),
		models.NewValueName(43, "Alligator"),
		models.NewValueName(44, "Splasher"),
		models.NewValueName(45, "Baron Warrior"),
		models.NewValueName(46, "Captain"),
		models.NewValueName(47, "Soul"),
		models.NewValueName(48, "Bloodbones"),
		models.NewValueName(49, "Ghoul"),
		models.NewValueName(50, "Skuldier"),
		models.NewValueName(51, "Revenant"),
		models.NewValueName(52, "Draculady"),
		models.NewValueName(53, "Cave Naga"),
		models.NewValueName(54, "Bone Dragon"),
		models.NewValueName(55, "Electrofish"),
		models.NewValueName(56, "Gigas Gator"),
		models.NewValueName(57, "Hydra"),
		models.NewValueName(58, "Flood Worm"),
		models.NewValueName(59, "Baron Guard"),
		models.NewValueName(60, "Lamia"),
		models.NewValueName(61, "Hell Needle"),
		models.NewValueName(62, "Twin Snake"),
		models.NewValueName(63, "Grudger"),
		models.NewValueName(64, "Mind Flayer"),
		models.NewValueName(65, "Naga"),
		models.NewValueName(66, "Ogre"),
		models.NewValueName(67, "Cait Sith"),
		models.NewValueName(68, "Black Knight"),
		models.NewValueName(69, "Centaur Knight"),
		models.NewValueName(70, "Miss Vamp"),
		models.NewValueName(71, "Puppeteer"),
		models.NewValueName(72, "Puppet"),
		models.NewValueName(73, "Gloomwing"),
		models.NewValueName(74, "Death Shell"),
		models.NewValueName(75, "Screamer"),
		models.NewValueName(76, "Warrior"),
		models.NewValueName(77, "Armadillo"),
		models.NewValueName(78, "Soldieress"),
		models.NewValueName(79, "Ghost Knight"),
		models.NewValueName(80, "King Naga"),
		models.NewValueName(81, "Medusa"),
		models.NewValueName(82, "Dinozombie"),
		models.NewValueName(83, "Magma Tortoise"),
		models.NewValueName(85, "Evil Dreamer"),
		models.NewValueName(86, "Bomb"),
		models.NewValueName(87, "Gray Bomb"),
		models.NewValueName(88, "Chimera"),
		models.NewValueName(89, "Hell Turtle"),
		models.NewValueName(90, "Ironback"),
		models.NewValueName(91, "Fiery Hound"),
		models.NewValueName(92, "Gorgon"),
		models.NewValueName(93, "Stone Golem"),
		models.NewValueName(94, "Lilith"),
		models.NewValueName(95, "Lamia Matriarch"),
		models.NewValueName(96, "Tiny Toad"),
		models.NewValueName(97, "Mors"),
		models.NewValueName(98, "Roc"),
		models.NewValueName(99, "Sorcerer"),
		models.NewValueName(100, "Mad Ogre"),
		models.NewValueName(101, "Evil Doll"),
		models.NewValueName(102, "Bloody Bat"),
		models.NewValueName(103, "Arachne"),
		models.NewValueName(104, "Fiery Knight"),
		models.NewValueName(105, "Beamer"),
		models.NewValueName(106, "Balloon"),
		models.NewValueName(107, "Dark Grenade"),
		models.NewValueName(108, "Plague"),
		models.NewValueName(109, "Last Arm"),
		models.NewValueName(110, "Coeurl"),
		models.NewValueName(111, "Mech Dragon"),
		models.NewValueName(112, "Tarantula"),
		models.NewValueName(113, "Gremlin"),
		models.NewValueName(114, "Undergrounder"),
		models.NewValueName(115, "Abyss Worm"),
		models.NewValueName(116, "Bloody Eye"),
		models.NewValueName(117, "Crawler"),
		models.NewValueName(118, "Ice Lizard"),
		models.NewValueName(119, "Hell Flapper"),
		models.NewValueName(120, "Evil Bat"),
		models.NewValueName(121, "Cold Beast"),
		models.NewValueName(122, "Summoner"),
		models.NewValueName(123, "Sorceress"),
		models.NewValueName(124, "Trap Door"),
		models.NewValueName(125, "Bog Witch"),
		models.NewValueName(126, "Mammon"),
		models.NewValueName(127, "Chimera Brain"),
		models.NewValueName(128, "Lunar Virus"),
		models.NewValueName(129, "Zemus's Breath"),
		models.NewValueName(130, "White Mousse"),
		models.NewValueName(131, "Black Flan"),
		models.NewValueName(132, "Zemus's Mind"),
		models.NewValueName(133, "Mythril Golem"),
		models.NewValueName(134, "Green Dragon"),
		models.NewValueName(135, "Flan Princess"),
		models.NewValueName(136, "Security Eye"),
		models.NewValueName(137, "Mech Soldier"),
		models.NewValueName(138, "Giant Soldier"),
		models.NewValueName(139, "Selene Guardian"),
		models.NewValueName(140, "Malboro"),
		models.NewValueName(141, "Silver Dragon"),
		models.NewValueName(142, "Yellow Dragon"),
		models.NewValueName(143, "Mini Satana"),
		models.NewValueName(144, "Eukaryote"),
		models.NewValueName(145, "Wicked Mask"),
		models.NewValueName(146, "Centaurion"),
		models.NewValueName(147, "Giant Warrior"),
		models.NewValueName(148, "Steel Golem"),
		models.NewValueName(149, "Ahriman"),
		models.NewValueName(150, "Lunasaur"),
		models.NewValueName(151, "Searcher"),
		models.NewValueName(152, "Dark Sage"),
		models.NewValueName(153, "Dark Bahamut"),
		models.NewValueName(154, "Prokaryote"),
		models.NewValueName(155, "Ogopogo"),
		models.NewValueName(156, "Blue Dragon"),
		models.NewValueName(157, "Gold Dragon"),
		models.NewValueName(158, "Thunder Dragon"),
		models.NewValueName(159, "White Dragon"),
		models.NewValueName(160, "Red Dragon"),
		models.NewValueName(161, "Behemoth"),
		models.NewValueName(162, "Mist Dragon"),
		models.NewValueName(163, "Octomammoth"),
		models.NewValueName(164, "Antlion"),
		models.NewValueName(165, "Mom Bomb"),
		models.NewValueName(166, "Scarmiglione"),
		models.NewValueName(167, "Scarmiglione"),
		models.NewValueName(168, "Baigan"),
		models.NewValueName(169, "Right Arm"),
		models.NewValueName(170, "Left Arm"),
		models.NewValueName(171, "Cagnazzo"),
		models.NewValueName(172, "Dark Elf"),
		models.NewValueName(173, "Dark Elf"),
		models.NewValueName(174, "Sandy"),
		models.NewValueName(175, "Cindy"),
		models.NewValueName(176, "Mindy"),
		models.NewValueName(177, "Golbez"),
		models.NewValueName(178, "Barbariccia"),
		models.NewValueName(179, "Calco"),
		models.NewValueName(180, "Calcobrena"),
		models.NewValueName(181, "Golbez"),
		models.NewValueName(182, "Shadow Dragon"),
		models.NewValueName(183, "Doctor"),
		models.NewValueName(184, "Barnabas"),
		models.NewValueName(185, "Dr. Lugae"),
		models.NewValueName(186, "King of Eblan"),
		models.NewValueName(187, "Queen of Eblan"),
		models.NewValueName(188, "Rubicante"),
		models.NewValueName(189, "Odin"),
		models.NewValueName(190, "Leviathan"),
		models.NewValueName(191, "Bahamut"),
		models.NewValueName(192, "Demon Wall"),
		models.NewValueName(193, "Asura"),
		models.NewValueName(194, "Elemental Lord"),
		models.NewValueName(195, "Elemental Lord"),
		models.NewValueName(196, "Dark Dragon"),
		models.NewValueName(197, "Sahagin"),
		models.NewValueName(198, "CPU"),
		models.NewValueName(199, "Defense Node"),
		models.NewValueName(200, "Zemus"),
		models.NewValueName(201, "Zeromus"),
		models.NewValueName(202, "Zeromus"),
		models.NewValueName(203, "Dragoon"),
		models.NewValueName(204, "Bard"),
		models.NewValueName(205, "Monk"),
		models.NewValueName(206, "Dark Knight"),
		models.NewValueName(207, "Girl"),
		models.NewValueName(209, "Floating Eyeball"),
		models.NewValueName(210, "Zu"),
		models.NewValueName(211, "Brina"),
		models.NewValueName(212, "Skullnant"),
		models.NewValueName(213, "Barnabas-Z"),
		models.NewValueName(214, "Attack Node"),
		models.NewValueName(217, "Zeromus"),
		models.NewValueName(224, "Mystery Egg"),
		models.NewValueName(225, "Dark Elf"),
		models.NewValueName(226, "Rubicante"),
		models.NewValueName(227, "Elemental Lord"),
		models.NewValueName(228, "Elemental Lord"),
		models.NewValueName(229, "Puppet"),
		models.NewValueName(230, "Evil Doll"),
		models.NewValueName(231, "Black Lizard"),
		models.NewValueName(232, "Lamia"),
		models.NewValueName(233, "King Naga"),
		models.NewValueName(234, "Green Dragon"),
		models.NewValueName(235, "Yellow Dragon"),
		models.NewValueName(236, "Fiery Hound"),
		models.NewValueName(237, "Coeurl"),
		models.NewValueName(238, "Ghost Knight"),
		models.NewValueName(239, "Lamia Matriarch"),
		models.NewValueName(240, "Mad Ogre"),
		models.NewValueName(241, "Green Dragon"),
		models.NewValueName(242, "Mythril Golem"),
		models.NewValueName(243, "Goblin"),
		models.NewValueName(244, "Hell Flapper"),
		models.NewValueName(245, "King Naga"),
		models.NewValueName(246, "Arachne"),
		models.NewValueName(247, "Thunder Dragon"),
		models.NewValueName(248, "Trap Door"),
		models.NewValueName(249, "Naga"),
		models.NewValueName(250, "Chimera"),
		models.NewValueName(251, "Fiery Hound"),
		models.NewValueName(252, "Stone Golem"),
		models.NewValueName(253, "Mech Soldier"),
		models.NewValueName(254, "Centaurion"),
		models.NewValueName(255, "Giant Soldier"),
		models.NewValueName(256, "Mech Dragon"),
		models.NewValueName(257, "Domovoi"),
		models.NewValueName(258, "Puppeteer"),
		models.NewValueName(259, "Sorcerer"),
		models.NewValueName(260, "Sorcerer"),
		models.NewValueName(261, "Sorcerer"),
		models.NewValueName(262, "Sorcerer"),
		models.NewValueName(263, "Sorcerer"),
		models.NewValueName(264, "Summoner"),
		models.NewValueName(265, "Summoner"),
		models.NewValueName(266, "Summoner"),
		models.NewValueName(267, "Summoner"),
		models.NewValueName(268, "Security Eye"),
		models.NewValueName(269, "Security Eye"),
		models.NewValueName(270, "Security Eye"),
		models.NewValueName(271, "Searcher"),
		models.NewValueName(272, "Searcher"),
		models.NewValueName(273, "Searcher"),
		models.NewValueName(274, "Golbez"),
	}
)
