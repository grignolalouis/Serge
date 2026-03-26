# Guide complet — La Supply Chain de Siam Mango Co.

## Vue d'ensemble

Siam Mango Co. est un **distributeur B2B de mangues** basé à Bangkok. L'entreprise achète des mangues (fraiches et transformées) auprès de fermes thaïlandaises, les stocke dans un entrepôt, et les livre à des clients professionnels (hôtels, restaurants, supermarchés).

```
FOURNISSEURS          ENTREPÔT              CLIENTS
(Fermes thaï)         (Bangkok)             (B2B Bangkok)

┌─────────────┐       ┌──────────────┐      ┌───────────────────┐
│ SUP-001     │──PO──►│              │──OR──►│ Hotels            │
│ SUP-002     │──PO──►│   WH-BKK     │──OR──►│ Restaurants       │
│ SUP-003     │──PO──►│   Bangkok    │──OR──►│ Supermarchés      │
│ SUP-004     │──PO──►│   Central    │──OR──►│ Marchés           │
│ SUP-005     │──PO──►│              │      └───────────────────┘
└─────────────┘       └──────┬───────┘
                             │
                      ┌──────▼───────┐
                      │  CARRIERS    │
                      │  Kerry       │
                      │  SCG Cold    │
                      │  Flash Exp.  │
                      └──────────────┘
```

La chaîne se lit de gauche à droite :
1. On **commande** aux fournisseurs (Purchase Order)
2. Les produits arrivent à l'**entrepôt** (inbound)
3. Un client passe une **commande** (Order)
4. On prépare et **expédie** via un transporteur (Shipment)
5. Le client **reçoit** sa livraison

---

## Les 4 systèmes

Notre supply chain est découpée en **4 systèmes** indépendants, chacun gérant une partie du flux :

| Système | Domaine métier | Ce qu'il gère | Analogie |
|---------|---------------|---------------|----------|
| **SRM** | Approvisionnement (Procurement) | Fournisseurs, commandes d'achat | "À qui on achète" |
| **WMS** | Stock (Inventory) | Produits, entrepôt, niveaux de stock, mouvements | "Ce qu'on a en stock" |
| **OMS** | Ventes (Sales) | Clients, commandes de vente | "Ce que les clients commandent" |
| **TMS** | Livraison (Shipping) | Transporteurs, expéditions, tracking | "Comment on livre" |

**Pourquoi 4 systèmes séparés ?** En entreprise réelle, chaque système est souvent un logiciel différent (SAP pour l'ERP, un WMS dédié, un TMS séparé). Les données sont **cloisonnées**. C'est exactement le problème que SERGE résout : un agent qui requête les 4 et fait le lien.

---

## Domaine 1 : Approvisionnement (SRM)

### Supplier (Fournisseur)

Un fournisseur est une ferme ou un producteur qui nous vend des mangues.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant unique | `SUP-001` |
| `name` | string | Nom de l'entreprise | "Somchai Mango Farm" |
| `region` | string | Région de Thaïlande où se situe la ferme | "Chanthaburi" |
| `contact` | string | Email de contact | "somchai@mangomail.th" |
| `lead_time_days` | int | Délai de livraison déclaré (en jours) | 3 |
| `rating` | float | Note de fiabilité (1.0 à 5.0) | 4.5 |
| `product_categories` | []string | Catégories de produits qu'il fournit | ["premium", "standard"] |
| `payment_terms` | string | Conditions de paiement | "net_30" (paiement à 30 jours) |

**Nos 5 fournisseurs :**
- **SUP-001 Somchai Mango Farm** (Chanthaburi) — Le meilleur. Rating 4.5, 3 jours de lead time, 100% on-time. Fournit les mangues premium et standard.
- **SUP-002 Northern Fruits Co.** (Chiang Mai) — Le moins fiable. Rating 3.2, 5 jours déclarés mais 8 en réalité, 0% on-time. Toujours en retard.
- **SUP-003 Prachuap Harvest** (Prachuap Khiri Khan) — Variable. Rating 2.8, 50% on-time. Pas cher mais imprévisible.
- **SUP-004 Rayong Tropical** (Rayong) — Bon fournisseur de produits transformés (mangue séchée, purée). Rating 4.0, fiable.
- **SUP-005 Isaan Organic Farms** (Nakhon Ratchasima) — Bio, lead time plus long (6 jours), rating 3.5. Fournit les produits premium et assortiments.

**À retenir :** Le `lead_time_days` est ce que le fournisseur *promet*. Le délai *réel* se calcule à partir des Purchase Orders (date de commande → date de réception effective). C'est cette différence qui permet de mesurer la fiabilité.

---

### Purchase Order (Commande d'achat)

Un Purchase Order (PO) est une commande que **nous passons à un fournisseur** pour réapprovisionner notre stock.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant unique | `PO-001` |
| `supplier_id` | string | Fournisseur concerné | `SUP-001` |
| `status` | enum | État de la commande | "received" |
| `order_date` | date | Date à laquelle on a passé la commande | 2026-01-02 |
| `expected_delivery` | date | Date de livraison attendue | 2026-01-05 |
| `actual_delivery` | date? | Date de livraison réelle (null si pas encore reçu) | 2026-01-05 |
| `lines` | []PurchaseOrderLine | Produits commandés (quantité, coût) | voir ci-dessous |

**Statuts possibles :**
```
pending → confirmed → shipped → received
                                   ↘ cancelled
```
- **pending** : on a passé la commande, le fournisseur ne l'a pas encore confirmée
- **confirmed** : le fournisseur a accepté
- **shipped** : le fournisseur a expédié les produits vers notre entrepôt
- **received** : on a reçu les produits à l'entrepôt (stock mis à jour)
- **cancelled** : commande annulée

**PurchaseOrderLine** (ligne de commande) :

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `product_id` | string | Produit commandé | `PROD-001` |
| `quantity` | int | Quantité | 100 |
| `unit_cost` | float | Coût unitaire d'achat (THB) | 280.00 |

**Point clé :** Le `unit_cost` est le prix auquel on **achète** au fournisseur. Le `unit_price` du produit est le prix auquel on **vend** au client. La différence est notre marge. Ex : on achète le Nam Doc Mai à 280 THB et on le vend à 450 THB → marge brute de ~60%.

**Méthodes métier :**
- `TotalCost()` : somme de (quantité × coût unitaire) pour toutes les lignes
- `IsOverdue()` : vrai si le PO n'est ni reçu ni annulé et que la date attendue est dépassée

---

## Domaine 2 : Stock (WMS)

### Product (Produit)

Un produit est un article que nous vendons. Chez Siam Mango Co., ce sont des mangues sous différentes formes et conditionnements.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant unique | `PROD-001` |
| `sku` | string | Code article (Stock Keeping Unit) | "MANGO-NDM-5" |
| `name` | string | Nom complet | "Nam Doc Mai Mango Box 5kg" |
| `category` | string | Catégorie | "premium" |
| `unit_price` | float | Prix de vente unitaire (THB) | 450.00 |
| `weight_kg` | float | Poids par unité | 5.0 |
| `reorder_point` | int | Seuil de réapprovisionnement | 20 |
| `shelf_life_days` | int | Durée de conservation (jours) | 7 |
| `storage_type` | string | Type de stockage requis | "refrigerated" |
| `supplier_ids` | []string | Fournisseurs qui proposent ce produit | ["SUP-001", "SUP-004"] |

**Nos 8 produits :**

| ID | Nom | Catégorie | Prix | Poids | Conservation | Stockage |
|----|-----|-----------|------|-------|-------------|----------|
| PROD-001 | Nam Doc Mai Mango Box 5kg | premium | 450฿ | 5kg | 7j | frigo |
| PROD-002 | Ok Rong Mango Box 10kg | standard | 650฿ | 10kg | 7j | frigo |
| PROD-003 | Keaw Mango Box 5kg | standard | 350฿ | 5kg | 10j | frigo |
| PROD-004 | Mahachanok Mango Box 10kg | premium | 800฿ | 10kg | 5j | frigo |
| PROD-005 | Mixed Mango Assortment 5kg | assorted | 500฿ | 5kg | 7j | frigo |
| PROD-006 | Dried Mango Slices 500g | processed | 180฿ | 0.5kg | 180j | ambiant |
| PROD-007 | Mango Puree 1L | processed | 220฿ | 1.1kg | 90j | frigo |
| PROD-008 | Mango Sticky Rice Kit 2-person | assorted | 320฿ | 1.5kg | 3j | frigo |

**À retenir :**
- Le `reorder_point` est le seuil sous lequel il faut passer un nouveau PO. Si le stock descend en dessous, c'est une alerte.
- Le `shelf_life_days` est critique pour les mangues fraiches (5-10 jours). Les produits transformés durent beaucoup plus longtemps (90-180 jours).
- Le `storage_type` détermine le choix du transporteur : un produit "refrigerated" **doit** voyager en camion frigorifique, pas en transport classique. C'est exactement ce qui a mal tourné avec SHP-009 (Kerry Express ground pour des mangues fraiches → panne du frigo).

---

### Warehouse (Entrepôt)

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant | `WH-BKK` |
| `name` | string | Nom | "Bangkok Central Warehouse" |
| `location` | string | Adresse | "Bang Na, Bangkok" |

Nous avons **un seul entrepôt** à Bangkok (Bang Na). Tous les produits y sont stockés, toutes les expéditions en partent.

---

### InventoryRecord (Niveau de stock)

Un enregistrement de stock représente la **quantité actuelle** d'un produit dans un entrepôt.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `product_id` | string | Produit | `PROD-001` |
| `warehouse_id` | string | Entrepôt | `WH-BKK` |
| `quantity` | int | Quantité totale en stock | 8 |
| `reserved` | int | Quantité réservée (pour des commandes en préparation) | 0 |
| `last_updated` | date | Dernière mise à jour | 2026-02-10 |

**Méthodes métier :**
- `Available()` = `quantity - reserved` → stock réellement disponible
- `NeedsReorder(reorderPoint)` → vrai si `Available() < reorderPoint`

**Exemple concret :** PROD-001 a 8 unités en stock, 0 réservées → 8 disponibles. Le reorder point est 20. Donc `NeedsReorder(20)` = vrai → **alerte stock bas**.

C'est le scénario clé du UC-2 : la commande ORD-007 a besoin de 50 unités mais il n'y en a que 8.

---

### StockMovement (Mouvement de stock)

Un mouvement de stock trace chaque entrée ou sortie de produit dans l'entrepôt.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant | `SM-001` |
| `product_id` | string | Produit concerné | `PROD-001` |
| `warehouse_id` | string | Entrepôt | `WH-BKK` |
| `type` | enum | Direction du mouvement | "inbound" ou "outbound" |
| `quantity` | int | Quantité déplacée | 100 |
| `reference_id` | string | ID de la source (PO ou Order) | `PO-001` |
| `reference_type` | string | Type de source | "purchase_order" ou "sales_order" |
| `created_at` | date | Quand le mouvement a eu lieu | 2026-01-05 |

**Deux types de mouvements :**
- **inbound** : produit qui ENTRE dans l'entrepôt ← vient d'un Purchase Order (le fournisseur livre)
- **outbound** : produit qui SORT de l'entrepôt → part pour une commande client (sales order)

**Le lien clé :** Le `reference_id` + `reference_type` permettent de tracer l'origine de chaque mouvement. On peut remonter de "le stock a baissé de 10" à "c'est parce que la commande ORD-001 a été expédiée".

---

## Domaine 3 : Ventes (OMS)

### Customer (Client)

Un client est une entreprise B2B qui nous achète des mangues.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant | `CUST-001` |
| `name` | string | Nom du contact | "Somying Prasertsak" |
| `company` | string | Nom de l'entreprise | "Siam Paragon Food Hall" |
| `segment` | enum | Type de client | "retail" ou "wholesale" |
| `region` | string | Région | "Bangkok" |
| `address` | string | Adresse de livraison | "991 Rama I Rd, Pathum Wan" |
| `delivery_zone` | string | Zone de livraison | "central_bkk" ou "outer_bkk" |

**Segments :**
- **retail** : supermarchés, food halls (Siam Paragon, Tops, Villa Market, Gourmet Market, Makro, Lotus's). Commandes moyennes, fréquentes.
- **wholesale** : hôtels et restaurants (Blue Elephant, Mandarin Oriental, Sukhothai, Gaggan, Banyan Tree, Chatuchak Market). Commandes plus grosses, souvent des standing orders hebdomadaires.

**Zones de livraison :**
- **central_bkk** : centre de Bangkok (Sathorn, Sukhumvit, Pathum Wan). Livraison rapide, 1 jour.
- **outer_bkk** : banlieue (Bangkapi, Rangsit, Chatuchak). Peut prendre 2 jours.

---

### Order (Commande de vente)

Une commande est une demande d'un client pour des produits.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant | `ORD-007` |
| `customer_id` | string | Client | `CUST-001` |
| `status` | enum | État de la commande | "processing" |
| `order_date` | date | Date de la commande | 2026-02-05 |
| `required_date` | date | Date limite demandée par le client | 2026-02-12 |
| `shipped_date` | date? | Date d'expédition (null si pas encore expédié) | null |
| `delivered_date` | date? | Date de livraison (null si pas encore livré) | null |
| `priority` | enum | Niveau d'urgence | "urgent" |
| `notes` | string? | Instructions spéciales | "Deliver before 9am" |
| `lines` | []OrderLine | Produits commandés | voir ci-dessous |

**Statuts possibles :**
```
pending → processing → shipped → delivered
                                    ↘ cancelled
```
- **pending** : commande reçue, pas encore traitée
- **processing** : en préparation (picking, packing). Si le stock est insuffisant, la commande reste bloquée ici.
- **shipped** : expédiée, un Shipment existe dans le TMS
- **delivered** : livrée au client
- **cancelled** : annulée

**OrderLine** (ligne de commande) :

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `product_id` | string | Produit commandé | `PROD-001` |
| `quantity` | int | Quantité | 50 |
| `unit_price` | float | Prix de vente unitaire (THB) | 450.00 |

**Méthodes métier :**
- `TotalAmount()` : somme de (quantité × prix unitaire) pour toutes les lignes
- `IsOverdue()` : vrai si la commande n'est ni livrée ni annulée et que la `required_date` est dépassée

**Le scénario clé :** ORD-007 est "processing" + "urgent" + overdue. L'agent doit comprendre POURQUOI en allant chercher dans le WMS (pas assez de stock) puis dans le SRM (un PO est en route mais pas encore arrivé).

---

## Domaine 4 : Livraison (TMS)

### Carrier (Transporteur)

Un carrier est une entreprise de transport qui livre nos produits aux clients.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant | `CAR-002` |
| `name` | string | Nom | "SCG Cold Chain" |
| `type` | string | Type de véhicule | "refrigerated" |
| `cost_per_kg` | float | Coût au kilo (THB) | 25.00 |
| `avg_transit_days` | int | Durée de transit moyenne | 1 |

**Nos 3 transporteurs :**

| ID | Nom | Type | Coût/kg | Transit | Usage |
|----|-----|------|---------|---------|-------|
| CAR-001 | Kerry Express | ground | 15฿ | 2 jours | Produits ambiants (mangue séchée) |
| CAR-002 | SCG Cold Chain | refrigerated | 25฿ | 1 jour | Produits frais — choix standard |
| CAR-003 | Flash Express Premium | refrigerated | 35฿ | 1 jour | Produits frais — premium, plus cher |

**Règle critique :** Un produit avec `storage_type: "refrigerated"` **doit** être transporté par un carrier `type: "refrigerated"`. Utiliser Kerry Express (ground) pour des mangues fraiches est une erreur — c'est ce qui a causé l'incident SHP-009 (panne de frigo sur un véhicule non adapté).

---

### Shipment (Expédition)

Un shipment représente l'envoi physique de produits d'un point A à un point B.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `id` | string | Identifiant | `SHP-005` |
| `order_id` | string | Commande associée | `ORD-005` |
| `carrier_id` | string | Transporteur | `CAR-002` |
| `status` | enum | État de l'expédition | "in_transit" |
| `tracking_number` | string | Numéro de suivi | "SCG-20260210-005" |
| `origin` | string | Point de départ | "Bangkok Central Warehouse" |
| `destination` | string | Point d'arrivée | "Villa Market, Bangkok" |
| `weight_kg` | float | Poids total | 200.0 |
| `shipped_at` | date? | Date d'expédition | 2026-02-10 08:00 |
| `estimated_arrival` | date? | Date d'arrivée estimée | 2026-02-12 17:00 |
| `delivered_at` | date? | Date de livraison effective | null (pas encore livré) |
| `exception_reason` | string? | Raison du problème (si status=exception) | "" |

**Statuts possibles :**
```
pending → in_transit → delivered
                ↘ exception
```
- **pending** : expédition créée mais pas encore partie
- **in_transit** : en route
- **delivered** : livrée
- **exception** : problème de livraison (panne, adresse incorrecte, refus, etc.)

**Méthode métier :**
- `IsLate()` : vrai si la date estimée est dépassée sans livraison, ou si la livraison a eu lieu après la date estimée

**Relation Order ↔ Shipment :** Normalement 1 commande = 1 expédition. Mais une commande peut être **splitée** en plusieurs shipments (ex : ORD-023 a 2 shipments — SHP-017 en frigorifique pour les mangues fraiches et SHP-018 en ground pour la mangue séchée).

---

### TrackingEvent (Événement de suivi)

Un tracking event est un point de passage dans la vie d'une expédition.

| Champ | Type | Description | Exemple |
|-------|------|-------------|---------|
| `shipment_id` | string | Expédition | `SHP-005` |
| `timestamp` | datetime | Quand | 2026-02-10 08:30 |
| `location` | string | Où | "Bang Na, Bangkok" |
| `status` | string | Statut à ce moment | "picked_up" |
| `description` | string | Description libre | "Package picked up from Bangkok Central Warehouse" |

**Séquence typique :**
1. `picked_up` — Colis récupéré à l'entrepôt
2. `in_transit` — En route (peut avoir plusieurs événements)
3. `delivered` — Livré au client

Si problème : un événement `exception` apparait avec la description du problème.

---

## Les liens entre systèmes

C'est ici que SERGE prend tout son sens. Les 4 systèmes sont connectés par des **références croisées** :

```
SRM                    WMS                    OMS                    TMS
───                    ───                    ───                    ───
Supplier               Product                Customer
    │                    │  ▲                    │
    ▼                    │  │                    ▼
PurchaseOrder ──────► StockMovement         Order ──────────────► Shipment
  (inbound)            (reference_id)        (order_id)            │
                         │                     │                    ▼
                         ▼                     ▼              TrackingEvent
                    InventoryRecord        OrderLine
                                          (product_id)
```

**Les jointures clés que l'agent doit faire :**

| Question | Chainage | Systèmes |
|----------|----------|----------|
| "Où est ma commande ?" | Order → Shipment → TrackingEvents | OMS → TMS |
| "Pourquoi la commande est bloquée ?" | Order → OrderLine → Inventory → PurchaseOrder | OMS → WMS → SRM |
| "Quel fournisseur est le meilleur ?" | Supplier → PurchaseOrders → dates | SRM |
| "Quel est l'historique de ce produit ?" | Product → StockMovements → POs + Orders | WMS → SRM + OMS |
| "Le bon transporteur a-t-il été choisi ?" | Shipment → Order → Product.storage_type vs Carrier.type | TMS → OMS → WMS |

---

## Exemple complet : pourquoi ORD-007 est bloquée

Voici le raisonnement cross-système que l'agent fait :

1. **OMS** → `oms_get_order("ORD-007")` : commande en "processing", urgente, besoin de 50x PROD-001 pour Siam Paragon, deadline Feb 12 dépassée
2. **TMS** → `tms_track_order("ORD-007")` : aucun shipment trouvé → pas encore expédiée
3. **WMS** → `wms_check_inventory("PROD-001")` : 8 unités disponibles, reorder point 20 → stock critique, impossible de remplir 50
4. **SRM** → `srm_list_purchase_orders_by_status("shipped")` : PO-008 de SUP-001, 150 unités de PROD-001 en route, ETA Feb 18

**Diagnostic :** La commande est bloquée parce que le stock de Nam Doc Mai est à 8 (besoin 50). Un réapprovisionnement (PO-008) arrive le 18 février, soit 6 jours après la deadline. La commande ne peut pas être honorée avant.

C'est exactement ce qu'un ops manager ferait manuellement en 15-30 minutes en ouvrant 4 logiciels différents. L'agent le fait en 4 appels en quelques secondes.

---

## Volumes de données

| Entité | Quantité | Période |
|--------|----------|---------|
| Fournisseurs | 5 | - |
| Produits | 8 | - |
| Entrepôts | 1 | - |
| Clients | 12 | - |
| Transporteurs | 3 | - |
| Commandes de vente | 500 | Oct 2025 – Feb 2026 |
| Commandes d'achat | 60 | Oct 2025 – Feb 2026 |
| Expéditions | 493 | Oct 2025 – Feb 2026 |
| Événements de suivi | 1 431 | Oct 2025 – Feb 2026 |
| Mouvements de stock | 965 | Oct 2025 – Feb 2026 |

Les 30 premières commandes (ORD-001 à ORD-030) et les 18 premiers POs (PO-001 à PO-018) sont les **scénarios d'évaluation** avec des situations précises (stockout, retard fournisseur, exception transport, etc.). Le reste est généré pour le réalisme du volume.
