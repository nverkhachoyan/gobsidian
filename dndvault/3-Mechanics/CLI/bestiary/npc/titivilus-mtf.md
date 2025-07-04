---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/16
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Titivilus"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 179
---
# [Titivilus](3-Mechanics\CLI\bestiary\npc/titivilus-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 179*  

> [!quote]- A quote from Mordenkainen  
> 
> I actually admire Titivilus. If he weren't so remorselessly evil, he'd be an excellent administrator of the Balance.

The gloomy Lord of the Second, Dispater, rules from his iron palace, seeming to hide surrounded by its labyrinthine corridors, iron walls, diabolical traps, and monstrous servants. So intense is his paranoia that he almost never travels farther than the sprawling city that lies outside his magnificent palace. Dispater knows he has enemies on all sides—enemies who would do to him what has been done to the likes of Geryon, Moloch, and so many others.

Dispater is correct to fear, but the true threat comes not from without. The lord's great error was allowing himself to be seduced by Titivilus, who beguiled his way into being the primary advisor in Dispater's household.

Although he is inferior in physical strength and power when compared to other archdevils, Titivilus compensates with cunning. A shrewd and calculating politician, he has clawed his way up through the ranks to become the second-most powerful fiend in Dis, entirely by saying just the right thing at the right time to get what he wanted. Charming and pleasant, he is a master at negotiation, able to twist words in such a way as to leave his victims confused and believing they have found a friend in Titivilus. Through these skills, Titivilus has manipulated everyone along his path to power, either to win them over to his cause or to remove them as a threat.

Since gaining his position, Titivilus has convinced Dispater that countless plots are being hatched against him and that Asmodeus himself seeks to remove Dispater from power. In response, Dispater has withdrawn to his palace and left day-to-day decisions to Titivilus, while also authorizing him to answer and negotiate bargains with mortals who attempt to summon Dispater. Titivilus now represents his master and speaks with his voice, a turn of events that leads some to whisper that either Titivilus is Dispater in disguise, or that Titivilus has removed the archduke and replaced him altogether.

Titivilus recognizes the inherent precariousness of his position. After all, Dispater's acceptance of his plans and his advice can last only so long before some other plotter steps in and reveals the truth. For insurance, Titivilus has begun recruiting outsiders to deal with problem devils, to insulate himself against criticism, and, above all, to create complications that he can solve so as to reinforce his value in the eyes of his master. Titivilus finds adventurers well suited to the tasks he needs performed and recruits them directly or through intermediaries, expending them later as his plans require.

```statblock
"name": "Titivilus (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "127"
"hit_dice": "17d8 + 51"
"stats":
- !!int "19"
- !!int "22"
- !!int "17"
- !!int "24"
- !!int "22"
- !!int "26"
"speed": "40 ft., fly 60 ft."
"saves":
  "Charisma": !!int "13"
  "Dexterity": !!int "11"
  "Wisdom": !!int "11"
  "Constitution": !!int "8"
"skillsaves":
  "Intimidation": !!int "13"
  "Deception": !!int "13"
  "Insight": !!int "11"
  "Persuasion": !!int "13"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "all, telepathy 120 ft."
"cr": "16"
"traits":
- "desc": "Titivilus's innate spellcasting ability is Charisma (spell save DC 21).\
    \ He can innately cast the following spells, requiring no material components:\n\
    \nAt will: [alter self](/3-Mechanics/CLI/spells/alter-self.md), [animate dead](/3-Mechanics/CLI/spells/animate-dead.md),\
    \ [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md), [confusion](/3-Mechanics/CLI/spells/confusion.md),\
    \ [major image](/3-Mechanics/CLI/spells/major-image.md), [modify memory](/3-Mechanics/CLI/spells/modify-memory.md),\
    \ [nondetection](/3-Mechanics/CLI/spells/nondetection.md), [sending](/3-Mechanics/CLI/spells/sending.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\n1/day each: [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md),\
    \ [symbol](/3-Mechanics/CLI/spells/symbol.md) (discord or sleep only)\n\n3/day\
    \ each: [greater invisibility](/3-Mechanics/CLI/spells/greater-invisibility.md)\
    \ (self only), [mislead](/3-Mechanics/CLI/spells/mislead.md)"
  "name": "Innate Spellcasting"
- "desc": "If Titivilus fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Titivilus has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Titivilus's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "Titivilus regains 10 hit points at the start of his turn. If he takes cold\
    \ or radiant damage, this trait doesn't function at the start of his next turn.\
    \ Titivilus dies only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
- "desc": "Whenever Titivilus speaks, he can choose a point within 60 feet; his voice\
    \ emanates from that point."
  "name": "Ventriloquism"
"actions":
- "desc": "Titivilus makes one sword attack and uses his Frightful Word once."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage plus dice:3d10|text(16)\
    \ (3d10) necrotic damage, or dice:1d10 + 4|text(9) (1d10 + 4) slashing damage\
    \ plus dice:3d10|text(16) (3d10) necrotic damage if used with two hands. If\
    \ the target is a creature, its hit point maximum is reduced by an amount equal\
    \ to half the necrotic damage it takes."
  "name": "Silver Sword"
- "desc": "Titivilus targets one creature he can see within 10 feet of him. The target\
    \ must succeed on a DC 21 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, the target must take the Dash action and move away from Titivilus\
    \ by the safest available route on each of its turns, unless there is nowhere\
    \ to move, in which case it needn't take the Dash action. The target can repeat\
    \ the saving throw at the end of each of its turns, ending the effect on itself\
    \ on a success."
  "name": "Frightful Word"
- "desc": "Titivilus magically teleports, along with any equipment he is wearing and\
    \ carrying, up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
- "desc": "Titivilus targets one creature he can see within 60 feet of him. The target\
    \ must make a DC 21 Charisma saving throw. On a failure the target is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ for 1 minute. The [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) target\
    \ can repeat the saving throw if Titivilus deals any damage to it. A creature\
    \ that succeeds on the saving throw is immune to Titivilus's Twisting Words for\
    \ 24 hours."
  "name": "Twisting Words"
"legendary_actions":
- "desc": "Titivilus attacks with his silver sword or uses his Frightful Word."
  "name": "Assault (Costs 2 Actions)"
- "desc": "Titivilus uses Twisting Words. Alternatively, he targets one creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by him that is within 60 feet of him; that [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ target must make a DC 21 Charisma saving throw. On a failure, Titivilus decides\
    \ how the target acts during its next turn."
  "name": "Corrupting Guidance"
- "desc": "Titivilus uses his Teleport action."
  "name": "Teleport"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Titivilus.webp"
```
^statblock

```encounter-table
name: Titivilus
creatures:
 - 1: Titivilus
```