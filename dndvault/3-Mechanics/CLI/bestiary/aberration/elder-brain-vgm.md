---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/14
- monster/environment/underdark
- monster/size/large
- monster/type/aberration
aliases: ["Elder Brain"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 173
---
# [Elder Brain](3-Mechanics\CLI\bestiary\aberration/elder-brain-vgm.md)
*Source: Volo's Guide to Monsters p. 173*  

The ultimate expression of illithid domination, an elder brain sprawls within a vat of viscous brine, touching the thoughts of creatures near and far. It scrawls upon the canvas of their minds, rewriting their thoughts and authoring their dreams.

## Psychic Infiltrators

When an elder brain infiltrates a mind, it alters the creature's perception and deceives its senses, causing it to see, hear, touch, taste, or feel reality according to the elder brain's intent. From across great distances, it implants subconscious suggestions or subtly influences dreams to compel creatures toward a course of action that benefits its grand plan.

When its insidious suggestions fail to take hold, an elder brain asserts its dominance more directly. It seizes control of a resistant mind and controls the creature's body as it would a puppet. Against the rare, strong-willed stalwart that defies it or attacks it, an elder brain sends a blast of overwhelming psychic force to crush the upstart's mind, rendering the creature a thoughtless, drooling shell.

When a mind flayer perishes, the elder brain's servants feed the contents of its skull to their master, which absorbs the illithid's brain and all the knowledge and experience contained therein. In this way the elder brain continually increases its knowledge, uniting the thoughts and experiences of the illithid colony into a unified whole. Mind flayers conceive of this "oneness" as a sacred state in the same way that a worshiper of a human deity might view an eternal afterlife in the heavens-for an elder brain can evoke the persona of any illithid it has ever absorbed.

The ambitions of an elder brain are always tempered by its relative immobility. Although its telepathic senses can reach for miles, moving anywhere is always a dangerous proposition. If forced outside its brine pool, an elder brain will swiftly expire, and transporting an elder brain in its pool through confining and tortuous subterranean tunnels frequently proves difficult or impossible.

## Devourer of Thoughts

An elder brain sustains itself by consuming the brains of other creatures. When the mind flayer servants that guard and tend to an elder brain don't bring its meals directly to it, the elder brain reaches out with tendrils of thought, mentally compelling creatures to come to it so that it may feed upon them.

## Hive Mind

Non-illithids call this creature an elder brain because it acts as the central communication hub for an entire mind flayer colony just as a brain does for a living body. Linked to the elder brain, the colony acts like a single organism, acting in concert as if each illithid were the digit of a hand.

## Ego Unhindered

Each elder brain considers itself and its desires the most important things in the multiverse, the mind flayers in its colony nothing more than extensions of its will. But no two elder brains are alike, and each presides over its colony according to its own unique personality and storehouse of collected knowledge and experience. Some elder brains reign as domineering tyrants, while others serve more benignly as sages, counselors, and repositories of information and lore for the mind flayers that protect and nourish them.

## An Elder Brain's Lair

The lair of an elder brain always lies deep in the heart of a mind flayer colony. The creature dwells in a dimly glowing brine pool, filled with foul and brackish water infused with the elder brain's vital fluids and with psionic energy.

```statblock
"name": "Elder Brain (VGM)"
"size": "Large"
"type": "aberration"
"alignment": "Lawful Evil"
"ac": !!int "10"
"hp": !!int "210"
"hit_dice": "20d10 + 100"
"stats":
- !!int "15"
- !!int "10"
- !!int "20"
- !!int "21"
- !!int "19"
- !!int "24"
"speed": "5 ft., swim 10 ft."
"saves":
  "Charisma": !!int "12"
  "Wisdom": !!int "9"
  "Intelligence": !!int "10"
"skillsaves":
  "Intimidation": !!int "12"
  "Deception": !!int "12"
  "Insight": !!int "14"
  "Arcana": !!int "10"
  "Persuasion": !!int "12"
"senses": "blindsight 120 ft., passive Perception 14"
"languages": "understands Common, Deep Speech, and Undercommon but can't speak, telepathy\
  \ 5 miles"
"cr": "14"
"traits":
- "desc": "The elder brain's innate spellcasting ability is Intelligence (spell save\
    \ DC 18). It can innately cast the following spells, requiring no components:\n\
    \nAt will: [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md)\n\n1/day each: [dominate\
    \ monster](/3-Mechanics/CLI/spells/dominate-monster.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\
    \ (self only)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "The elder brain is aware of the presence of creatures within 5 miles of\
    \ it that have an Intelligence score of 4 or higher. It knows the distance and\
    \ direction to each creature, as well as each one's intelligence score, but can't\
    \ sense anything else about it. A creature protected by a [mind blank](/3-Mechanics/CLI/spells/mind-blank.md)\
    \ spell, a [nondetection](/3-Mechanics/CLI/spells/nondetection.md) spell, or similar\
    \ magic can't be perceived in this manner."
  "name": "Creature Sense"
- "desc": "If the elder brain fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The elder brain has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "The elder brain can use its telepathy to initiate and maintain telepathic\
    \ conversations with up to ten creatures at a time. The elder brain can let those\
    \ creatures telepathically hear each other while connected in this way."
  "name": "Telepathic Hub"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 30 ft., one target.\
    \ Hit: dice:4d8 + 2|text(20) (4d8 + 2) bludgeoning damage. If the target\
    \ is a Huge or smaller creature, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 15) and takes dice:1d8 + 5|text(9) (1d8 + 5) psychic damage at\
    \ the start of each of its turns until the grapple ends. The elder brain can have\
    \ up to four targets [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ at a time."
  "name": "Tentacle"
- "desc": "The elder brain magically emits psychic energy. Creatures of the elder\
    \ brain's choice within 60 feet of it must succeed on a DC 18 Intelligence saving\
    \ throw or take dice:5d10 + 5|text(32) (5d10 + 5) psychic damage and be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for 1 minute. A target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Mind Blast (Recharge 5-6)"
- "desc": "The elder brain targets one [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ creature it can perceive with its Creature Sense trait and establishes a psychic\
    \ link with that creature. Until the psychic link ends, the elder brain can perceive\
    \ everything the target senses. The target becomes aware that something is linked\
    \ to its mind once it is no longer [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ and the elder brain can terminate the link at any time (no action required).\
    \ The target can use an action on its turn to attempt to break the psychic link,\
    \ doing so with a successful DC 18 Charisma saving throw. On a successful save,\
    \ the target takes dice:3d6|text(10) (3d6) psychic damage. The psychic link\
    \ also ends if the target and the elder brain are more than 5 miles apart, with\
    \ no consequences to the target. The elder brain can form psychic links with up\
    \ to ten creatures at a time."
  "name": "Psychic Link"
- "desc": "The elder brain targets a creature with which it has a psychic link. The\
    \ elder brain gains insight into the target's reasoning, its emotional state,\
    \ and thoughts that loom large in its mind (including things the target worries\
    \ about, loves, or hates). The elder brain can also make a Charisma ([Deception](/3-Mechanics/CLI/rules/skills.md#Deception))\
    \ check with advantage to deceive the target's mind into thinking it believes\
    \ one idea or feels a particular emotion. The target contests this attempt with\
    \ a Wisdom ([Insight](/3-Mechanics/CLI/rules/skills.md#Insight)) check. If the\
    \ elder brain succeeds, the mind believes the deception for 1 hour or until evidence\
    \ of the lie is presented to the target."
  "name": "Sense Thoughts"
"legendary_actions":
- "desc": "The elder brain makes a tentacle attack."
  "name": "Tentacle"
- "desc": "The elder brain targets a creature within 120 feet of it with which it\
    \ has a psychic link. The elder brain breaks the creature's [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ on a spell it has cast. The creature also takes dice: 1d4|avg|noform (1d4)\
    \ psychic damage per level of the spell."
  "name": "Break Concentration"
- "desc": "The elder brain targets a creature within 120 feet of it with which it\
    \ has a psychic link. Enemies of the elder brain within 10 feet of that creature\
    \ take dice:3d6|text(10) (3d6) psychic damage."
  "name": "Psychic Pulse"
- "desc": "The elder brain targets a creature within 120 feet of it with which it\
    \ has a psychic link. The elder brain ends the link, causing the creature to have\
    \ disadvantage on all ability checks, attack rolls, and saving throws until the\
    \ end of the creature's next turn."
  "name": "Sever Psychic Link"
"lair_actions":
- "desc": "When fighting inside its lair, an elder brain can use lair actions. On\
    \ initiative count 20 (losing initiative ties), an elder brain can take one lair\
    \ action to cause one of the following effects; the elder brain can't use the\
    \ same lair action two rounds in a row:"
  "name": ""
- "desc": "- The elder brain casts [wall of force](/3-Mechanics/CLI/spells/wall-of-force.md).\
    \  \n- The elder brain targets one friendly creature it can sense within 120 feet\
    \ of it. The target has a flash of inspiration and gains advantage on one attack\
    \ roll, ability check, or saving throw it makes before the end of its next turn.\
    \ If the target doesn't or can't use this benefit in that time, the inspiration\
    \ is lost.  \n- The elder brain targets one creature it can sense within 120 feet\
    \ of it and anchors it by sheer force of will. The target must succeed on a DC\
    \ 18 Charisma saving throw or be unable to leave its current space. It can repeat\
    \ the saving throw at the end of each of its turns, ending the effect on itself\
    \ on a success.  "
  "name": ""
"regional_effects":
- "desc": "The territory within 5 miles of an elder brain is altered by the creature's\
    \ psionic presence, which creates one or more of the following effects:"
  "name": ""
- "desc": "- Creatures within 5 miles of an elder brain feel as if they are being\
    \ followed, even when they are not.  \n- The elder brain can overhear any telepathic\
    \ conversation happening within 5 miles of it. The creature that initiated the\
    \ telepathic conversation makes a DC 18 Wisdom ([Insight](/3-Mechanics/CLI/rules/skills.md#Insight))\
    \ check when telepathic contact is first established. If the check succeeds, the\
    \ creature is aware that something is eavesdropping on the conversation. The nature\
    \ of the eavesdropper isn't revealed, and the elder brain can't participate in\
    \ the telepathic conversation unless it has formed a psychic link with the creature\
    \ that initiated it.  \n- Any creature with which the elder brain has formed a\
    \ psychic link hears faint, incomprehensible whispers in the deepest recesses\
    \ of its mind. This psychic detritus consists of the elder brain's stray thoughts\
    \ commingled with those of other creatures to which it is linked.  "
  "name": ""
- "desc": "If the elder brain dies, these effects immediately end."
  "name": ""
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Elder%20Brain.webp"
```
^statblock

```encounter-table
name: Elder Brain
creatures:
 - 1: Elder Brain
```

## Environment

underdark