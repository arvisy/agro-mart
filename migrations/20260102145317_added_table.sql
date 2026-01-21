-- 
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
    CREATE TYPE user_role AS ENUM ('farmer', 'seller', 'admin');
  END IF;
END$$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
    CREATE TYPE order_status AS ENUM ('PENDING', 'PROCESSING', 'COMPLETED', 'FAILED');
  END IF;
END$$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status') THEN
    CREATE TYPE payment_status AS ENUM ('SUCCESS', 'FAILED');
  END IF;
END$$;

-- 
CREATE TABLE IF NOT EXISTS user (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  role user_role NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- 
CREATE TABLE IF NOT EXISTS product_category (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL
);

-- 
CREATE TABLE IF NOT EXISTS product (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  category_id INT NOT NULL,
  name VARCHAR(150) NOT NULL,
  unit VARCHAR(20),         -- kg, liter, pack
  price NUMERIC(14,2) NOT NULL,
  stock INT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now(),

  CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES product_categories(id)
);

-- 
CREATE TABLE IF NOT EXISTS order (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  status order_status NOT NULL,
  total_amount NUMERIC(14,2) NOT NULL,
  idempotency_key VARCHAR(100) UNIQUE,
  created_at TIMESTAMPTZ DEFAULT now(),

  CONSTRAINT fk_orders_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 
CREATE TABLE IF NOT EXISTS order_item (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_id UUID NOT NULL,
  product_id UUID NOT NULL,
  quantity INT NOT NULL,
  price NUMERIC(14,2) NOT NULL,

  CONSTRAINT fk_order_items_order FOREIGN KEY (order_id) REFERENCES orders(id),
  CONSTRAINT fk_order_items_product FOREIGN KEY (product_id) REFERENCES products(id)
);

-- 
CREATE TABLE IF NOT EXISTS payment (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_id UUID NOT NULL,
  method VARCHAR(50),
  status payment_status NOT NULL,
  paid_at TIMESTAMPTZ,

  CONSTRAINT fk_payments_order FOREIGN KEY (order_id) REFERENCES orders(id)
);
