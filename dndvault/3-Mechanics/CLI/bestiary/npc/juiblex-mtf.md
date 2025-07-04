---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/23
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Juiblex"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 151
---
# [Juiblex](3-Mechanics\CLI\bestiary\npc/juiblex-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 151*  

## Juiblex

Called the Faceless Lord and the Oozing Hunger in ancient grimoires, Juiblex is demon lord of slime and ooze, a noxious creature that doesn't care about the plots and schemes of others of its kind. It exists only to consume, digesting and transforming living matter into more of itself.

A true horror, Juiblex is a mass of bubbling slime, swirling black and green, with glaring red eyes floating and shifting within it. It can rise up like a 20-foot hill, lashing out with dripping pseudopods to drag victims into its bulk. Those consumed by Juiblex are obliterated.

Only the truly insane worship Juiblex and tend to its slimes and oozes. Those who offer themselves up to the demon lord are engulfed by it and become vaguely humanoid, sentient oozes. The bodies of these former flesh-and-blood creatures form Juiblex's extended physical body, while the demon lord slowly digests and savors their identities over time.

## Juiblex's Lair

Juiblex's principal lair is known as the Slime Pits, a realm which Juiblex shares with [Zuggtmoy](/3-Mechanics/CLI/bestiary/npc/zuggtmoy-mtf.md). This layer of the Abyss, which is also known as Shedaklah, is a bubbling morass of oozing, fetid sludge. Its landscape is covered in vast expanses of caustic and unintelligent slimes, and strange organic forms rise from the oceans of molds and oozes at Juiblex's command.

```statblock
"name": "Juiblex (MTF)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "350"
"hit_dice": "28d12 + 168"
"stats":
- !!int "24"
- !!int "10"
- !!int "23"
- !!int "20"
- !!int "20"
- !!int "16"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "7"
  "Wisdom": !!int "12"
  "Constitution": !!int "13"
"skillsaves":
  "Perception": !!int "12"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "truesight 120 ft., passive Perception 22"
"languages": "all, telepathy 120 ft."
"cr": "23"
"traits":
- "desc": "Juiblex's spellcasting ability is Charisma (spell save DC 18, dice: d20+10\
    \ (+10) to hit with spell attacks). Juiblex can innately cast the following\
    \ spells, requiring no material components:\n\nAt will: [acid splash](/3-Mechanics/CLI/spells/acid-splash.md)\
    \ (17th level), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\n3/day\
    \ each: [blight](/3-Mechanics/CLI/spells/blight.md), [contagion](/3-Mechanics/CLI/spells/contagion.md),\
    \ [gaseous form](/3-Mechanics/CLI/spells/gaseous-form.md)"
  "name": "Innate Spellcasting"
- "desc": "Any creature, other than an ooze, that starts its turn within 10 feet of\
    \ Juiblex must succeed on a DC 21 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the start of the creature's next turn."
  "name": "Foul"
- "desc": "If Juiblex fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Juiblex has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Juiblex's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "Juiblex regains 20 hit points at the start of its turn. If it takes fire\
    \ or radiant damage, this trait doesn't function at the start of its next turn.\
    \ Juiblex dies only if it starts its turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
- "desc": "Juiblex can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "Juiblex makes three acid lash attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+14 (+14) to hit, reach 10 ft., one\
    \ target. Hit: dice:4d6 + 7|text(21) (4d6 + 7) acid damage. Any creature\
    \ killed by this attack is drawn into Juiblex's body, and the corpse is obliterated\
    \ after 1 minute."
  "name": "Acid Lash"
- "desc": "Juiblex spews out a corrosive slime, targeting one creature that it can\
    \ see within 60 feet of it. The target must make a DC 21 Dexterity saving throw.\
    \ On a failure, the target takes dice:10d10|text(55) (10d10) acid damage.\
    \ Unless the target avoids taking any of this damage, any metal armor worn by\
    \ the target takes a permanent −1 penalty to the AC it offers, and any metal weapon\
    \ it is carrying or wearing takes a permanent −1 penalty to damage rolls. The\
    \ penalty worsens each time a target is subjected to this effect. If the penalty\
    \ on an object drops to −5, the object is destroyed."
  "name": "Eject Slime (Recharge 5-6)"
"legendary_actions":
- "desc": "Juiblex casts [acid splash](/3-Mechanics/CLI/spells/acid-splash.md)."
  "name": "Acid Splash"
- "desc": "Juiblex makes one acid lash attack."
  "name": "Attack"
- "desc": "Melee Weapon Attack: dice: d20+14 (+14) to hit, reach 10 ft., one\
    \ creature. Hit: dice:4d6 + 7|text(21) (4d6 + 7) poison damage, and the\
    \ target is slimed. Until the slime is scraped off with an action, the target\
    \ is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), and any creature,\
    \ other than an ooze, is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ while within 10 feet of the target."
  "name": "Corrupting Touch (Costs 2 Actions)"
"lair_actions":
- "desc": "On Initiative count 20 (losing initiative ties), Juiblex can take a lair\
    \ action to cause one of the following effects; he can't use the same effect two\
    \ rounds in a row:"
  "name": ""
- "desc": "- Juiblex slimes a square area of ground it can see within the lair. The\
    \ area can be up to 10 feet on a side. The slime lasts for 1 hour or until it\
    \ is burned away with fire. When the slime appears, each creature in that area\
    \ must succeed on a DC 21 Strength saving throw or become [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ When a creature enters the area for the first time on a turn or ends its turn\
    \ there, that creature must make the same save. A [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ creature is stuck as long as it remains in the slimy area or until it breaks\
    \ free. The [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) creature,\
    \ or another creature who can reach it, can use its action to try to break free\
    \ and must succeed on a DC 21 Strength check. If the slime is set on fire, it\
    \ burns away after 1 round. Any creature that starts its turn in the burning slime\
    \ takes dice:4d10|text(22) (4d10) fire damage.  \n- Juiblex slimes a square\
    \ area of ground it can see within the lair. The area can be up to 10 feet on\
    \ a side. The slime lasts for 1 hour or until it is burned away with fire. When\
    \ the slime appears, each creature on it must succeed on a DC 21 Dexterity saving\
    \ throw or fall [prone](/3-Mechanics/CLI/rules/conditions.md#prone) and slide\
    \ 10 feet in a random direction determined by a dice: d8|avg|noform (d8) roll.\
    \ When a creature enters the area for the first time on a turn or ends its turn\
    \ there, that creature must make the same save. If the slime is set on fire, it\
    \ burns away after 1 round. Any creature that starts its turn in the burning slime\
    \ takes dice:4d10|text(22) (4d10) fire damage.  \n- A [green slime](/3-Mechanics/CLI/traps-hazards/green-slime.md)\
    \ (see the Dungeon Master's Guide) appears on a spot on the ceiling that Juiblex\
    \ chooses within the lair. The slime disintegrates after 1 hour.  "
  "name": ""
"regional_effects":
- "desc": "The region containing Juiblex's lair is warped by his magic, creating one\
    \ or more of the following effects:"
  "name": ""
- "desc": "- Small bodies of water, such as ponds or wells, within 1 mile of the lair\
    \ turn highly acidic, corroding any object that touches them. Surfaces within\
    \ 6 miles of the lair are frequently covered by a thin film of slime, which is\
    \ slick and sticks to anything that touches it.  \n- If a humanoid spends at least\
    \ 1 hour within 1 mile of the lair, that creature must succeed on a DC 18 Wisdom\
    \ saving throw or descend into a madness determined by the Madness of Juiblex\
    \ table. A creature that succeeds on this saving throw can't be affected by this\
    \ regional effect again for 24 hours.  "
  "name": ""
- "desc": "If Juiblex dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
- "desc": "If a creature goes mad in Juiblex's lair or within line of sight of the\
    \ demon lord, roll on the Madness of Juiblex table to determine the nature of\
    \ the madness, which is a character flaw that lasts until cured. See the \"Dungeon\
    \ Master's Guide\" for more on madness.\n\nMadness of Juiblex\n\ndice: [](juiblex-mtf.md#^madness-of-juiblex)\n\
    \n| dice: d100 | Flaw (lasts until cured) |\n|------------|--------------------------|\n\
    | 01–20 | \"I must consume everything I can!\" |\n| 21–40 | \"I refuse to part\
    \ with any of my possessions.\" |\n| 41–60 | \"I'll do everything I can to get\
    \ others to eat and drink beyond their normal limits.\" |\n| 61–80 | \"I must\
    \ possess as many material goods as I can.\" |\n| 81–00 | \"My personality is\
    \ irrelevant. I am defined by what I consume.\" |\n^madness-of-juiblex"
  "name": "Madness of Juiblex"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Juiblex.webp"
```
^statblock

```encounter-table
name: Juiblex
creatures:
 - 1: Juiblex
```