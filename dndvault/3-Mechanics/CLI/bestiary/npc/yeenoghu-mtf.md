---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/24
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Yeenoghu"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 155
---
# [Yeenoghu](3-Mechanics\CLI\bestiary\npc/yeenoghu-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 155*  

## Yeenoghu

The Beast of Butchery appears as a great battle-scarred gnoll, towering 14 feet tall. Yeenoghu is the Gnoll Lord, and his creations are made in his twisted image. When the demon lord hunted across the Material Plane, packs of hyenas followed in his wake. Those that ate of great Yeenoghu's kills became gnolls, emulating their master's ways. Few others worship the Beast of Butchery, but those who do tend to take on a gnoll-like aspect, hunched over, and filing their teeth down to points.

Yeenoghu wants nothing more than slaughter and senseless destruction. The gnolls are his instruments, and he drives them to ever-greater atrocities in his name. Yeenoghu takes pleasure in causing fear before death, and he sows sorrow and despair through destroying beloved things. He doesn't parlay; to meet him is to do battle with him—unless he becomes bored. The Beast of Butchery has a long rivalry with Baphomet, the Horned King, and the two demon lords and their followers attack one another on sight.

The Gnoll Lord is covered in matted fur and taut, leathery hide, his face like a grinning predator's skull. Patchwork armor made of discarded shields and breastplates is lashed onto his body with heavy chains, decorated by the flayed skins of his foes. He wields a triple-headed flail called the Butcher, which he can summon into his hand at will, although he is as likely to tear his prey apart with his bare hands before ripping out its throat with his teeth.

## Yeenoghu's Lair

Yeenoghu's lair in the Abyss is called the Death Dells, its barren hills and ravines serving as one great hunting ground, where he pursues captured mortals in a cruel game. Yeenoghu's lair is a place of blood and death, populated by [gnolls](/3-Mechanics/CLI/bestiary/humanoid/gnoll.md), [hyenas](/3-Mechanics/CLI/bestiary/beast/hyena.md), and [ghouls](/3-Mechanics/CLI/bestiary/undead/ghoul.md), and there are few structures or signs of civilization on his layer of the Abyss.

```statblock
"name": "Yeenoghu (MTF)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "333"
"hit_dice": "23d12 + 184"
"stats":
- !!int "29"
- !!int "16"
- !!int "26"
- !!int "15"
- !!int "24"
- !!int "15"
"speed": "50 ft."
"saves":
  "Dexterity": !!int "10"
  "Wisdom": !!int "14"
  "Constitution": !!int "15"
"skillsaves":
  "Intimidation": !!int "9"
  "Perception": !!int "14"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 24"
"languages": "all, telepathy 120 ft."
"cr": "24"
"traits":
- "desc": "Yeenoghu's spellcasting ability is Charisma (spell save DC 17, dice: d20+9\
    \ (+9) to hit with spell attacks). He can innately cast the following spells,\
    \ requiring no material components:\n\nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\
    \n1/day: [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\n3/day each:\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fear](/3-Mechanics/CLI/spells/fear.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md)"
  "name": "Innate Spellcasting"
- "desc": "If Yeenoghu fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Yeenoghu has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Yeenoghu's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "When Yeenoghu reduces a creature to 0 hit points with a melee attack on\
    \ his turn, Yeenoghu can take a bonus action to move up to half his speed and\
    \ make a bite attack."
  "name": "Rampage"
"actions":
- "desc": "Yeenoghu makes three flail attacks. If an attack hits, he can cause it\
    \ to create an additional effect of his choice or at random (each effect can be\
    \ used only once per Multiattack):\n\n1. The attack deals an extra dice:2d12|text(13)\
    \ (2d12) bludgeoning damage.\n\n2. The target must succeed on a DC 17 Constitution\
    \ saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ until the start of Yeenoghu's next turn.\n\n3. The target must succeed on a\
    \ DC 17 Wisdom saving throw or be affected by the [confusion](/3-Mechanics/CLI/spells/confusion.md)\
    \ spell until the start of Yeenoghu's next turn."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 15 ft., one\
    \ target. Hit: dice:1d12 + 9|text(15) (1d12 + 9) bludgeoning damage."
  "name": "Flail"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:1d10 + 9|text(14) (1d10 + 9) piercing damage."
  "name": "Bite"
"legendary_actions":
- "desc": "Yeenoghu moves up to his speed."
  "name": "Charge"
- "desc": "Yeenoghu makes a flail attack. If the attack hits, the target must succeed\
    \ on a DC 24 Strength saving throw or be pushed 15 feet in a straight line away\
    \ from Yeenoghu. If the saving throw fails by 5 or more, the target falls [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Swat Away"
- "desc": "Yeenoghu makes a bite attack against each creature within 10 feet of him."
  "name": "Savage (Costs 2 Actions)"
"lair_actions":
- "desc": "On Initiative count 20 (losing initiative ties), Yeenoghu can take a lair\
    \ action to cause one of the following effects; he can't use the same effect two\
    \ rounds in a row:"
  "name": ""
- "desc": "- Yeenoghu causes an iron spike - 5 feet tall and 1 inch in diameter -\
    \ to burst from the ground at a point he can see within 100 feet of him. Any creature\
    \ in the space where the spike emerges must make a DC 24 Dexterity saving throw.\
    \ On a failed save, the creature takes dice:6d8|text(27) (6d8) piercing damage\
    \ and is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) by being\
    \ impaled on the spike. A creature can use an action to remove itself (or a creature\
    \ it can reach) from the spike, ending the [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ condition.  \n- Each gnoll or hyena that Yeenoghu can see can use its reaction\
    \ to move up to its speed.  \n- Until the next initiative count 20, all gnolls\
    \ and hyenas within the lair are enraged, causing them to have advantage on melee\
    \ weapon attack rolls and causing attack rolls to have advantage against them.\
    \  "
  "name": ""
"regional_effects":
- "desc": "The region containing Yeenoghu's lair is warped by his magic, creating\
    \ one or more of the following effects:"
  "name": ""
- "desc": "- Within 1 mile of the lair, large iron spikes grow out of the ground and\
    \ stone surfaces. Yeenoghu impales the bodies of the slain on these spikes.  \n\
    - Predatory beasts within 6 miles of the lair become unusually savage, killing\
    \ far more than what they need for food. Carcasses of prey are left to rot in\
    \ an unnatural display of wasteful slaughter.  \n- If a humanoid spends at least\
    \ 1 hour within 1 mile of the lair, that creature must succeed on a DC 17 Wisdom\
    \ saving throw or descend into a madness determined by the Madness of Yeenoghu\
    \ table. A creature that succeeds on this saving throw can't be affected by this\
    \ regional effect again for 24 hours.  "
  "name": ""
- "desc": "If Yeenoghu dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
- "desc": "If a creature goes mad in Yeenoghu's lair or within line of sight of the\
    \ demon lord, roll on the Madness of Yeenoghu table to determine the nature of\
    \ the madness, which is a character flaw that lasts until cured. See the \"Dungeon\
    \ Master's Guide\" for more on madness.\n\nMadness of Yeenoghu\n\ndice: [](yeenoghu-mtf.md#^madness-of-yeenoghu)\n\
    \n| dice: d100 | Flaw (lasts until cured) |\n|------------|--------------------------|\n\
    | 01–20 | \"I get caught up in the flow of anger, and try to stoke others around\
    \ me into forming an angry mob.\" |\n| 21–40 | \"The flesh of other intelligent\
    \ creatures is delicious!\" |\n| 41–60 | \"I rail against the laws and customs\
    \ of civilization, attempting to return to a more primitive time.\" |\n| 61–80\
    \ | \"I hunger for the deaths of others, and am constantly starting fights in\
    \ the hope of seeing bloodshed.\" |\n| 81–00 | \"I keep trophies from the bodies\
    \ I have slain, turning them into adornments.\" |\n^madness-of-yeenoghu"
  "name": "Madness of Yeenoghu"
"source":
- "MTF"
- "OotA"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Yeenoghu.webp"
```
^statblock

```encounter-table
name: Yeenoghu
creatures:
 - 1: Yeenoghu
```