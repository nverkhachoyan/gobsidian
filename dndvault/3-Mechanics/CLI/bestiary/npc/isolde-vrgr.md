---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/5
- monster/size/medium
- monster/type/fey/elf
aliases: ["Isolde"]
NoteIcon: monster
BestiaryType: fey (elf)
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 86
---
# [Isolde](3-Mechanics\CLI\bestiary\npc/isolde-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 86*  

Isolde was a holy warrior devoted to a pantheon of elven deities called the Seldarine. In this role, she defended the Feywild against dragons, demons, and other threats. In time, her heroics caught the eye of an archfey named Zybilna, who had forged secret pacts with some of the fiends Isolde and her companions had slain. Rather than be angry at Isolde, Zybilna was impressed by her. She enlisted a powerful fiend known only as "the Caller" to corrupt and slay all of Isolde's companions, leaving Isolde alone, bitter, and vulnerable. The insidious archfey then befriended Isolde and offered to help her forget her terrible losses. Isolde became the master of a traveling fey carnival that served as a gateway to Zybilna's domain. The carnival did what Zybilna hoped it would do: it brought comfort to Isolde and quelled her thirst for vengeance.

Zybilna and Isolde enjoyed a strong partnership for years, but as time wore on, they grew distant until their relationship finally soured. Eladrin crave change, yet Isolde felt like she was frozen in time. She wished to leave the fey carnival and pursue other dreams, but Zybilna wouldn't hear of it and secretly used [wish](/3-Mechanics/CLI/spells/wish.md) spells to make Isolde place her devotion to the carnival above her desire to leave it.

When Isolde's fey carnival crossed paths with another carnival from the Shadowfell, the eladrin found the escape she longed for. Isolde orchestrated a trade with the other carnival's owners, a pair of shadar-kai (elves native to the Shadowfell). Isolde would become the master of their carnival, and they would become the masters of hers. To appease Zybilna, this arrangement would remain in place only until the two carnivals crossed paths again.

Zybilna was intrigued enough by the shadar-kai to let Isolde go, but not without casting a spell that made Isolde forget about Zybilna and her Feywild domain, thus preventing the eladrin from divulging the archfey's secrets. As a further punishment, Zybilna sent malevolent fey creatures to hound Isolde and her Shadowfell carnival. Isolde doesn't know who is behind this petty torment, nor does she care. Her hunt for the Caller and her thirst for vengeance have become all-consuming.

```statblock
"name": "Isolde (VRGR)"
"size": "Medium"
"type": "fey"
"subtype": "elf"
"alignment": "Any Evil alignment"
"ac": !!int "19"
"ac_class": "[scale mail](/3-Mechanics/CLI/items/scale-mail.md)"
"hp": !!int "82"
"hit_dice": "11d8 + 33"
"stats":
- !!int "18"
- !!int "18"
- !!int "16"
- !!int "14"
- !!int "12"
- !!int "16"
"speed": "30 ft., fly 60 ft."
"saves":
  "Charisma": !!int "6"
  "Intelligence": !!int "5"
  "Strength": !!int "7"
  "Constitution": !!int "6"
"skillsaves":
  "Intimidation": !!int "6"
  "Deception": !!int "6"
  "Stealth": !!int "7"
  "Perception": !!int "4"
"damage_resistances": "cold; fire; lightning; poison; bludgeoning, piercing, slashing\
  \ from nonmagical attacks"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Abyssal, Common, Infernal"
"cr": "5"
"traits":
- "desc": "Isolde's spellcasting ability is Charisma (spell save DC 14). Isolde can\
    \ innately cast the following spells, requiring no material components:\n\n1/day:\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md) (self only)\n\n3/day\
    \ each: [alter self](/3-Mechanics/CLI/spells/alter-self.md), [command](/3-Mechanics/CLI/spells/command.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)"
  "name": "Innate Spellcasting"
- "desc": "The AC of Isolde includes her Charisma bonus."
  "name": "Fiendish Blessing"
- "desc": "While holding [Nepenthe](/3-Mechanics/CLI/items/nepenthe-vrgr.md), Isolde\
    \ creates an aura in a 10-foot radius around her. While this aura is active, Isolde\
    \ and all creatures friendly to her in the aura have advantage on saving throws\
    \ against spells and other magical effects."
  "name": "Magic Resistance Aura"
"actions":
- "desc": "Isolde makes two melee attacks or uses its Fire Ray twice."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft.., one\
    \ target. Hit: dice:1d8 + 7|text(11) (1d8 + 7) slashing damage, or dice:1d10\
    \ + 7|text(12) (1d10 + 7) slashing damage if used with two hands to make a\
    \ melee attack. If the target is a fiend or an undead, it takes an extra dice:2d10|text(11)\
    \ (2d10) radiant damage."
  "name": "Nepenthe"
- "desc": "Ranged Spell Attack: dice: d20+7 (+7) to hit, range 120 ft., one\
    \ target. Hit: dice:3d6|text(10) (3d6) fire damage."
  "name": "Fire Ray"
- "desc": "One humanoid Isolde can see within 30 feet of it must succeed on a DC 14\
    \ Wisdom saving throw or be magically [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ for 1 day. The [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) target\
    \ obeys Isolde's spoken commands. If the target suffers any harm from Isolde or\
    \ another creature or receives a suicidal command from Isolde, the target can\
    \ repeat the saving throw, ending the effect on itself on a success. If a target's\
    \ saving throw is successful, or if the effect ends for it, the creature is immune\
    \ to Isolde's Fiendish Charm for the next 24 hours."
  "name": "Fiendish Charm"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Isolde.webp"
```
^statblock

```encounter-table
name: Isolde
creatures:
 - 1: Isolde
```