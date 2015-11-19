# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20151119042840) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "debts", force: :cascade do |t|
    t.integer  "amount_in_cents"
    t.integer  "person_id"
    t.integer  "loaner_id"
    t.integer  "expense_id"
    t.boolean  "paid"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "debts", ["expense_id"], name: "index_debts_on_expense_id", using: :btree
  add_index "debts", ["loaner_id"], name: "index_debts_on_loaner_id", using: :btree
  add_index "debts", ["person_id"], name: "index_debts_on_person_id", using: :btree

  create_table "delayed_jobs", force: :cascade do |t|
    t.integer  "priority",   default: 0
    t.integer  "attempts",   default: 0
    t.text     "handler"
    t.text     "last_error"
    t.datetime "run_at"
    t.datetime "locked_at"
    t.datetime "failed_at"
    t.text     "locked_by"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  create_table "expenses", force: :cascade do |t|
    t.integer  "amount_in_cents"
    t.integer  "house_id"
    t.string   "name"
    t.text     "description"
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "creator_id"
    t.text     "notes"
    t.integer  "loaner_id"
  end

  add_index "expenses", ["creator_id"], name: "index_expenses_on_creator_id", using: :btree
  add_index "expenses", ["house_id"], name: "index_expenses_on_house_id", using: :btree
  add_index "expenses", ["loaner_id"], name: "index_expenses_on_loaner_id", using: :btree

  create_table "expenses_people", id: false, force: :cascade do |t|
    t.integer "expense_id"
    t.integer "person_id"
  end

  add_index "expenses_people", ["expense_id"], name: "index_expenses_people_on_expense_id", using: :btree
  add_index "expenses_people", ["person_id"], name: "index_expenses_people_on_person_id", using: :btree

  create_table "houses", force: :cascade do |t|
    t.string   "name"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  create_table "payments", force: :cascade do |t|
    t.integer  "person_paid_id"
    t.integer  "person_paying_id"
    t.integer  "amount_in_cents"
    t.datetime "created_at"
    t.datetime "updated_at"
    t.text     "notes"
    t.integer  "house_id"
  end

  add_index "payments", ["house_id"], name: "index_payments_on_house_id", using: :btree
  add_index "payments", ["person_paid_id"], name: "index_payments_on_person_paid_id", using: :btree
  add_index "payments", ["person_paying_id"], name: "index_payments_on_person_paying_id", using: :btree

  create_table "people", force: :cascade do |t|
    t.string   "name"
    t.string   "email"
    t.datetime "created_at"
    t.datetime "updated_at"
    t.integer  "house_id"
    t.string   "crypted_password",          limit: 128, default: "", null: false
    t.string   "salt",                      limit: 128, default: "", null: false
    t.string   "remember_token",            limit: 40
    t.datetime "remember_token_expires_at"
    t.string   "persistence_token",                     default: "", null: false
    t.string   "perishable_token",                      default: "", null: false
    t.integer  "login_count",                           default: 0,  null: false
    t.integer  "failed_login_count",                    default: 0,  null: false
    t.datetime "last_request_at"
    t.datetime "current_login_at"
    t.datetime "last_login_at"
    t.string   "current_login_ip"
    t.string   "last_login_ip"
  end

  add_index "people", ["house_id"], name: "index_people_on_house_id", using: :btree

  create_table "signups", force: :cascade do |t|
    t.string   "email"
    t.string   "name"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  create_table "users", force: :cascade do |t|
    t.string  "first_name"
    t.string  "last_name"
    t.integer "age"
  end

end
