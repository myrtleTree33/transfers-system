-- CreateTable
CREATE TABLE "StripeCheckoutSession" (
    "email" TEXT NOT NULL,
    "stripe_checkout_session_id" TEXT NOT NULL,
    "status" TEXT NOT NULL DEFAULT 'pending',
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "StripeCheckoutSession_email_key" ON "StripeCheckoutSession"("email");

-- CreateIndex
CREATE UNIQUE INDEX "StripeCheckoutSession_stripe_checkout_session_id_key" ON "StripeCheckoutSession"("stripe_checkout_session_id");

-- CreateIndex
CREATE UNIQUE INDEX "StripeCheckoutSession_email_stripe_checkout_session_id_key" ON "StripeCheckoutSession"("email", "stripe_checkout_session_id");
