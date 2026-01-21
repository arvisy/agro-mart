-- 
CREATE TABLE IF NOT EXISTS "user" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  role user_role NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- 
CREATE TABLE IF NOT EXISTS product_category (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- 
CREATE TABLE IF NOT EXISTS "product" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  category_id INT NOT NULL,
  name VARCHAR(150) NOT NULL,
  unit VARCHAR(20),         -- kg, liter, pack
  price NUMERIC(14,2) NOT NULL,
  stock INT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now(),

  CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES product_category(id)
);

-- 
CREATE TABLE IF NOT EXISTS "order" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  status order_status NOT NULL,
  total_amount NUMERIC(14,2) NOT NULL,
  idempotency_key VARCHAR(100) UNIQUE,
  created_at TIMESTAMPTZ DEFAULT now(),

  CONSTRAINT fk_orders_user FOREIGN KEY (user_id) REFERENCES "user"(id)
);

-- 
CREATE TABLE IF NOT EXISTS "order_item" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_id UUID NOT NULL,
  product_id UUID NOT NULL,
  quantity INT NOT NULL,
  price NUMERIC(14,2) NOT NULL,

  CONSTRAINT fk_order_items_order FOREIGN KEY (order_id) REFERENCES "order"(id),
  CONSTRAINT fk_order_items_product FOREIGN KEY (product_id) REFERENCES "product"(id)
);

-- 
CREATE TABLE IF NOT EXISTS "payment" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_id UUID NOT NULL,
  method VARCHAR(50),
  status payment_status NOT NULL,
  paid_at TIMESTAMPTZ,

  CONSTRAINT fk_payments_order FOREIGN KEY (order_id) REFERENCES "order"(id)
);
