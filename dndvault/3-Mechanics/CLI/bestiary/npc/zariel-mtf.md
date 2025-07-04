---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/26
- monster/size/large
- monster/type/fiend/devil
aliases: ["Zariel"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 180
---
# [Zariel](3-Mechanics\CLI\bestiary\npc/zariel-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 180*  

> [!quote]- A quote from Mordenkainen  
> 
> Do not pity the fallen angel. Fallen angels survive the fall. How many other souls did Zariel bring down with her?

## Zariel

Zariel rules Avernus, the first layer of the Nine Hells. Once a mighty angel charged with watching the tides of the Blood War, she succumbed to the plane's corrupting influence and fell from grace. She recently reclaimed her position as archdevil of Avernus after the cautious Bel proved inadequate at marshaling his forces to launch offensives against the encroaching demons. Now Bel advises her and helps her manage the war, though many whisper that her true agenda is vengeance against Asmodeus, and her true plan is to drive him from the Nine Hells.

All who enter and exit the Nine Hells must pass through Avernus, so the infernal armies muster on this layer. Here, the amnizus guard the citadels overlooking the River Styx, much of the fighting of the Blood War takes place, and devils gather to invade the Abyss. Anyone hoping to reach the lower layers must first contend with the darkness of this layer and the myriad threats it houses. Zariel manages it all and has the ultimate say over who comes and goes.

Given her role in the Blood War, Zariel is keenly interested in collecting souls from the greatest warriors on the Material Plane and securing their loyalty. She bargains hard, and mortals end up worse for dealing with her, because she holds all the cards. A bargain with Zariel is eternal; there is little hope of wriggling out of it.

However, she does expect the best from her servants, and so she allows her mortal followers to live out their lives, ever honing their talents, so she can put them to the best use when she finally calls in their debts. As a result, Zariel's servants are universally effective, disciplined, and dangerous.

## Zariel's Lair

Zariel makes her lair in a basalt citadel that rises up in Avernus. From nearly a mile away, one can hear the screams and moans coming from the burned victims chained to the stronghold's wall, the dying remains of those who failed to impress the archdevil. The stronghold, covering five square miles, is surrounded by walls reinforced with high turrets. Devils of all kinds crawl over the structure, ensuring that no intruders breach their defenses.

```statblock
"name": "Zariel (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "21"
"ac_class": "natural armor"
"hp": !!int "580"
"hit_dice": "40d10 + 360"
"stats":
- !!int "27"
- !!int "24"
- !!int "28"
- !!int "26"
- !!int "27"
- !!int "30"
"speed": "50 ft., fly 150 ft."
"saves":
  "Charisma": !!int "18"
  "Wisdom": !!int "16"
  "Intelligence": !!int "16"
"skillsaves":
  "Intimidation": !!int "18"
  "Perception": !!int "16"
"damage_resistances": "cold; fire; radiant; bludgeoning, piercing, slashing from nonmagical\
  \ attacks that aren't silvered"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 26"
"languages": "all, telepathy 120 ft."
"cr": "26"
"traits":
- "desc": "Zariel's innate spellcasting ability is Charisma (spell save DC 26). She\
    \ can innately cast the following spells, requiring no material components:\n\n\
    At will: [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium\
    \ when changing her appearance), [detect evil and good](/3-Mechanics/CLI/spells/detect-evil-and-good.md),\
    \ [fireball](/3-Mechanics/CLI/spells/fireball.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only), [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)\n\n3/day\
    \ each: [blade barrier](/3-Mechanics/CLI/spells/blade-barrier.md), [dispel evil\
    \ and good](/3-Mechanics/CLI/spells/dispel-evil-and-good.md), [finger of death](/3-Mechanics/CLI/spells/finger-of-death.md)"
  "name": "Innate Spellcasting"
- "desc": "Magical darkness doesn't impede Zariel's darkvision."
  "name": "Devil's Sight"
- "desc": "Zariel's weapon attacks are magical. When she hits with any weapon, the\
    \ weapon deals an extra dice:8d8|text(36) (8d8) fire damage (included in the\
    \ weapon attacks below)."
  "name": "Magic Weapons"
- "desc": "If Zariel fails a saving throw, she can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Zariel has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Zariel regains 20 hit points at the start of her turn. If she takes radiant\
    \ damage, this trait doesn't function at the start of her next turn. Zariel dies\
    \ only if she starts her turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Zariel attacks twice with her longsword or with her javelins. She can substitute\
    \ Horrid Touch for one of these attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) slashing damage plus dice:8d8|text(36)\
    \ (8d8) fire damage, or dice:2d10 + 8|text(19) (2d10 + 8) slashing damage\
    \ plus dice:8d8|text(36) (8d8) fire damage if used with two hands."
  "name": "Longsword"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+16 (+16) to hit, reach 10\
    \ ft. or range 30/120 ft., one target. Hit: dice:2d6 + 8|text(15) (2d6 +\
    \ 8) piercing damage plus dice:8d8|text(36) (8d8) fire damage."
  "name": "Javelin"
- "desc": "Zariel touches one creature within 10 feet of her. The target must succeed\
    \ on a DC 26 Constitution saving throw or take dice:8d10|text(44) (8d10) necrotic\
    \ damage and be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for\
    \ 1 minute. While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in\
    \ this way, the target is also [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ and [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened). The target can\
    \ repeat the saving throw at the end of each of its turns, ending the effect on\
    \ itself on a success."
  "name": "Horrid Touch (Recharge 5-6)"
- "desc": "Zariel magically teleports, along with any equipment she is wearing and\
    \ carrying, up to 120 feet to an unoccupied space she can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Zariel turns her magical gaze toward one creature she can see within 120\
    \ feet of her and commands it to combust. The target must succeed on a DC 26 Wisdom\
    \ saving throw or take dice:4d10|text(22) (4d10) fire damage."
  "name": "Immolating Gaze (Costs 2 Actions)"
- "desc": "Zariel uses her Teleport action."
  "name": "Teleport"
"lair_actions":
- "desc": "On initiative count 20 (losing initiative ties), Zariel can take a lair\
    \ action to cause one of the following effects; she can't use the same effect\
    \ two rounds in a row:"
  "name": ""
- "desc": "- Zariel casts [major image](/3-Mechanics/CLI/spells/major-image.md) four\
    \ times at its lowest level, targeting different areas with the spell. Zariel\
    \ prefers to create images of intruders' loved ones being burned alive. Zariel\
    \ doesn't need to concentrate on the spells, which end on initiative count 20\
    \ of the next round. Each creature that can see these illusions must succeed on\
    \ a DC 26 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the illusion for 1 minute. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success.  \n- Zariel casts her innate [fireball](/3-Mechanics/CLI/spells/fireball.md)\
    \ spell.  "
  "name": ""
"regional_effects":
- "desc": "The region containing Zariel's lair is warped by her magic, which creates\
    \ one or more of the following effects:"
  "name": ""
- "desc": "- The area within 9 miles of the lair is filled with screaming voices and\
    \ the stench of burning meat.  \n- Once every 60 feet within 1 mile of the lair,\
    \ 10-foot-high gouts of flame rise from the ground. Any creature or object that\
    \ touches the flame takes dice:2d6|text(7) (2d6) fire damage, though it can\
    \ take this damage no more than once per round.  \n- The area within 2 miles,\
    \ but no closer than 500 feet, of the lair is filled with smoke, which causes\
    \ the area to be heavily obscured. The smoke can't be cleared away.  "
  "name": ""
- "desc": "If Zariel dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Zariel.webp"
```
^statblock

```encounter-table
name: Zariel
creatures:
 - 1: Zariel
```